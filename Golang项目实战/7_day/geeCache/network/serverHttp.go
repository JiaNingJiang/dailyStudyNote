package network

import (
	"fmt"
	"geeCache/consistenthash"
	"geeCache/geecache"
	pb "geeCache/geecachepb"
	"geeCache/peer"
	"net/http"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

type HTTPPool struct {
	self     string // 记录自己的IP:PORT地址
	basePath string // url路径前缀

	mu          sync.Mutex             // guards peers and httpGetters
	hashRing    *consistenthash.Map    // 类型是一致性哈希算法的 Map，用来根据具体的 key 选择节点。
	httpGetters map[string]*httpGetter // 映射远程节点IP与对应的 httpGetter。每一个远程节点对应一个 httpGetter
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	fmt.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

// ServeHTTP handle all http requests
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := geecache.GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 将从本地获取的资源Marshal为protobuff格式
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write(body)
}

// 传入真实节点的IP地址作为名称,构建一致性哈希环，同时为各真实节点绑定httpGetter(客户端程序,通过http通信向远程节点请求资源)
func (p *HTTPPool) Set(nodes ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.hashRing = consistenthash.NewHashRing(defaultReplicas, nil)
	p.hashRing.Add(nodes...)
	p.httpGetters = make(map[string]*httpGetter, len(nodes))
	for _, node := range nodes {
		p.httpGetters[node] = &httpGetter{baseURL: node + p.basePath}
	}
}

// 根据资源的key获取其在一致性哈希环上对应的真实节点的 httpGetter
func (p *HTTPPool) PickPeer(key string) (peer.PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.hashRing.Get(key); peer != "" && peer != p.self { //如果发现资源存储在当前节点，则不返回httpGetter，因为可以在本地直接查询
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}
