package registry

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type GeeRegistry struct {
	timeout time.Duration // 服务超时时间
	mu      sync.Mutex    // protect following
	servers map[string]*ServerItem
}

type ServerItem struct {
	Addr  string
	start time.Time // 服务注册的时间
}

const (
	defaultPath    = "/_geerpc_/registry"
	defaultTimeout = time.Minute * 5 // 默认超时时间设置为 5 min，也就是说，任何注册的服务超过 5 min，即视为不可用状态。
)

func New(timeout time.Duration) *GeeRegistry {
	return &GeeRegistry{
		servers: make(map[string]*ServerItem),
		timeout: timeout,
	}
}

var DefaultGeeRegister = New(defaultTimeout)

// 注册一个新的服务实例(包含该实例的地址)或者更新实例的start time
func (r *GeeRegistry) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[addr]
	if s == nil {
		r.servers[addr] = &ServerItem{Addr: addr, start: time.Now()}
	} else {
		s.start = time.Now() // if exists, update start time to keep alive
	}
}

// 返回所有仍有效的服务实例的addr(完成排序)，对于已超时的服务实例将其删除
func (r *GeeRegistry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) { // 未设置超时时间或者未超出超时时间
			alive = append(alive, addr)
		} else {
			delete(r.servers, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

// 注册中心路由的核心功能
func (r *GeeRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // 服务于客户端(返回可用服务实例addr)
		w.Header().Set("X-Geerpc-Servers", strings.Join(r.aliveServers(), ",")) // 返回所有有效服务实例的addr
	case "POST": // 服务于服务端(心跳更新)
		addr := req.Header.Get("X-Geerpc-Server")
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr) // 心跳服务，注册或更新服务实例
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// 服务中心注册路由服务
func (r *GeeRegistry) HandleHTTP(registryPath string) {
	http.Handle(registryPath, r)
	fmt.Println("rpc registry path:", registryPath)
}

func HandleHTTP() {
	DefaultGeeRegister.HandleHTTP(defaultPath)
}

// 服务端：用于向注册中心发送心跳包
func Heartbeat(registry, addr string, duration time.Duration) {
	if duration == 0 {
		// make sure there is enough time to send heart beat
		// before it's removed from registry
		duration = defaultTimeout - time.Duration(1)*time.Minute // 心跳包时间间隔必须要小于超时时间
	}
	var err error
	err = sendHeartbeat(registry, addr) // 向注册中心发送自己的addr
	go func() {
		t := time.NewTicker(duration)
		for err == nil { // 只要未发生错误,就会持续发送心跳包
			<-t.C
			err = sendHeartbeat(registry, addr)
		}
	}()
}

// 以POST报文向注册中心发送心跳包
func sendHeartbeat(registry, addr string) error {
	fmt.Println(addr, "send heart beat to registry", registry)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", registry, nil)
	req.Header.Set("X-Geerpc-Server", addr)
	if _, err := httpClient.Do(req); err != nil {
		fmt.Println("rpc server: heart beat err:", err)
		return err
	}
	return nil
}
