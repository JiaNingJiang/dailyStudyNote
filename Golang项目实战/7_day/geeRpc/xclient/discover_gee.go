package xclient

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GeeRegistryDiscovery struct {
	*MultiServersDiscovery               // 进行服务实例选择
	registry               string        // 注册中心addr
	timeout                time.Duration // 服务列表的过期时间
	lastUpdate             time.Time     // 最后从注册中心更新服务列表的时间，默认 10s 过期，即 10s 之后，需要从注册中心更新新的列表。
}

const defaultUpdateTimeout = time.Second * 10

func NewGeeRegistryDiscovery(registerAddr string, timeout time.Duration) *GeeRegistryDiscovery {
	if timeout == 0 {
		timeout = defaultUpdateTimeout
	}
	d := &GeeRegistryDiscovery{
		MultiServersDiscovery: NewMultiServerDiscovery(make([]string, 0)),
		registry:              registerAddr,
		timeout:               timeout,
	}
	return d
}

// 手动更新服务列表(更新d.MultiServersDiscovery.servers)
func (d *GeeRegistryDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers // (d.servers继承自d.MultiServersDiscovery)
	d.lastUpdate = time.Now()
	return nil
}

// 从注册中心重新申请获取服务列表(更新d.MultiServersDiscovery.servers)
func (d *GeeRegistryDiscovery) Refresh() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.lastUpdate.Add(d.timeout).After(time.Now()) { // 距离上次服务更新还未到达服务列表的过期时间
		return nil
	}
	fmt.Println("rpc registry: refresh servers from registry", d.registry)
	resp, err := http.Get(d.registry) // 向注册中心重新申请获取服务实例的addr
	if err != nil {
		fmt.Println("rpc registry refresh err:", err)
		return err
	}
	servers := strings.Split(resp.Header.Get("X-Geerpc-Servers"), ",")
	d.servers = make([]string, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now()
	return nil
}

// 从注册中心更新服务列表，依据负载均衡策略返回一个服务实例
func (d *GeeRegistryDiscovery) Get(mode SelectMode) (string, error) {
	if err := d.Refresh(); err != nil {
		return "", err
	}
	return d.MultiServersDiscovery.Get(mode)
}

// 从注册中心更新服务列表，并返回所有服务实例
func (d *GeeRegistryDiscovery) GetAll() ([]string, error) {
	if err := d.Refresh(); err != nil {
		return nil, err
	}
	return d.MultiServersDiscovery.GetAll()
}
