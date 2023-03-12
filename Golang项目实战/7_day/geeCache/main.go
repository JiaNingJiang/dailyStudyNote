package main

import (
	"flag"
	"fmt"
	"geeCache/geecache"
	"geeCache/network"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// 在本地创建scores资源Group，注册GetterFunc用本地模拟的db作为数据源
func createGroup() *geecache.Group {
	return geecache.NewGroup("scores", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, group *geecache.Group) {
	server := network.NewHTTPPool(addr) // 根据本机的IP地址创建http server(实现了peer.PeerPicker接口)
	server.Set(addrs...)                // 根据所有节点的IP地址构建一致性哈希环，为每一个节点绑定httpGetter(请求客户端)
	group.RegisterPeers(server)         //为本地的资源Group绑定peer.PeerPicker接口
	fmt.Println("geecache is running at", addr)
	fmt.Println(http.ListenAndServe(addr[7:], server)) // 启动http server,负责监听来自于其他远程节点的查询请求
}

func startAPIServer(apiAddr string, group *geecache.Group) {
	http.Handle("/api", http.HandlerFunc( // 为本地的http server注册路由方法
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key") // 获取用户请求的资源key值
			view, err := group.Get(key)     // 查询对应的数据(本地 or 远程, 取决于资源key)，返回数据的副本
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			response := make([]byte, 0)
			response = append(response, []byte("Server Response: Query resulted value --")...)
			response = append(response, view.ByteSlice()...)
			response = append(response, []byte("            ")...)
			w.Write(response) // 将数据value写入到响应报文中

		}))
	fmt.Println("fontend server is running at", apiAddr)
	fmt.Println(http.ListenAndServe(apiAddr[7:], nil)) // 启动http server,负责监听来自于用户的查询请求

}

func main() {
	var port int // 当前节点在集群中负责通信的端口号
	var api bool // 是否开启用于API模型(允许用户向本节点发起资源请求)
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999" // 用户进行资源请求的url
	addrMap := map[int]string{         // 分布式集群节点间用于通信的url
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	group := createGroup()
	if api {
		go startAPIServer(apiAddr, group)
	}
	startCacheServer(addrMap[port], []string(addrs), group)
}
