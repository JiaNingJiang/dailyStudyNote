package geecache

import (
	"fmt"
	pb "geeCache/geecachepb"
	"geeCache/peer"
	"geeCache/singleflight"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error) // 需用户自行设计实现，用于从数据源(非缓存)获取数据

func (gf GetterFunc) Get(key string) ([]byte, error) {
	return gf(key)
}

type Group struct {
	name      string // 缓存表的名称
	getter    Getter // 缓存表的Getter接口
	mainCache cache  // 缓存表的并发缓存对象(支持读写并发)

	picker peer.PeerPicker // 节点查询接口

	loader *singleflight.Group // 防止缓存击穿的资源请求器
}

var (
	mu     sync.RWMutex                 // 缓存池并发锁
	groups = make(map[string]*Group, 0) // 缓存池
)

// 为缓存表注册节点查询接口
func (g *Group) RegisterPeers(picker peer.PeerPicker) {
	if g.picker != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.picker = picker
}

// 创建新的缓存表,并存入缓存池中
func NewGroup(name string, cacheBytesLimit int, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	group := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytesLimit: cacheBytesLimit},
		loader:    &singleflight.Group{},
	}
	mu.Lock()
	defer mu.Unlock()

	groups[name] = group

	return group
}

// 从缓存池取出指定name的缓存表
func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()

	return groups[name]
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok { // 缓存命中
		fmt.Println("[GeeCache] hit")
		return v, nil
	}

	return g.load(key) // 缓存未命中，需要从数据源获取数据到缓存中
}

func (g *Group) load(key string) (ByteView, error) {

	view, err := g.loader.Do(key, func() (interface{}, error) {
		if g.picker != nil {
			if peer, ok := g.picker.PickPeer(key); ok { // 根据资源key锁定对应的peerGetter(只有是远程节点才会继续向下执行，若查询失败或为本地节点，则跳转到g.getLocally(key)执行)
				if value, err := g.getFromPeer(peer, key); err == nil {
					return value, nil
				} else {
					fmt.Println("[GeeCache] Failed to get from peer", err)
				}
			}
		}
		return g.getLocally(key) // 从本地数据源缓存数据到本地
	})

	return view.(ByteView), err
}

// 调用对应peer节点的Get方法,根据group和key获取资源value
func (g *Group) getFromPeer(picker peer.PeerGetter, key string) (ByteView, error) {

	req := &pb.Request{Group: g.name, Key: key}
	res := &pb.Response{}

	err := picker.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Value}, nil
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key) // 采用用户自定义方式从数据源获取数据
	if err != nil {
		return ByteView{}, fmt.Errorf("%w + ", err)
	}

	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value) // 进行缓存
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
