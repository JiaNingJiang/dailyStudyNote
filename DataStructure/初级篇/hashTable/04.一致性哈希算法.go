package hashTable

import (
	"crypto/sha1"
	"sync"

	//	"hash"
	"math"
	"sort"
	"strconv"
)

const (
	DefaultVirualSpots = 400 // 每个真实节点可以拥有的虚拟节点上限个数
)

type node struct {
	nodeKey   string // 真实节点的标识符
	spotValue uint32 // 虚拟节点的哈希值
}

type nodesArray []node // 根据虚拟节点的哈希值大小进行升序排序

func (p nodesArray) Len() int           { return len(p) }
func (p nodesArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }
func (p nodesArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p nodesArray) Sort()              { sort.Sort(p) }

// HashRing store nodes and weigths
type HashRing struct {
	virualSpots int
	nodes       nodesArray     // 虚拟节点集合(所有的虚拟节点,按照虚拟节点的哈希值进行排序)
	weights     map[string]int // 真实节点的标识符和权重(该节点的所有虚拟节点权重之和)
	mu          sync.RWMutex
}

// 新建单节点虚拟节点上限数为spots的哈希环
func NewHashRing(spots int) *HashRing {
	if spots == 0 {
		spots = DefaultVirualSpots
	}

	h := &HashRing{
		virualSpots: spots,
		weights:     make(map[string]int),
	}
	return h
}

// AddNodes add nodes to hash ring
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.generate()
}

// 如果节点不存在,在哈希环上新增一个节点
func (h *HashRing) AddNode(nodeKey string, weight int) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.weights[nodeKey]; !ok {
		h.weights[nodeKey] = weight
	} else {
		return false
	}
	h.generate()
	return true
}

// 将一个真实节点的哈希环上的权重更新为nil,然后重新生成哈希环(其实就是删除该真实节点)
func (h *HashRing) RemoveNode(nodeKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.weights, nodeKey)
	h.generate()
}

// 如果节点已经存在,更新该真实节点在哈希环上的权重，然后重新生成哈希环
func (h *HashRing) UpdateNode(nodeKey string, weight int) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if oldWeight, ok := h.weights[nodeKey]; ok { // 该节点必须已经存在于哈希环上,且更新前后权重不一样
		if oldWeight != weight {
			h.weights[nodeKey] = weight
		} else {
			return false
		}
	} else {
		return false
	}
	h.generate()
	return true
}

// 重新生成哈希环(初始化h.nodes)，生成所有的虚拟节点,存入h.nodes
func (h *HashRing) generate() {
	var totalW int // 哈希环的总权重
	for _, w := range h.weights {
		totalW += w
	}

	totalVirtualSpots := h.virualSpots * len(h.weights) // 哈希环上的总虚拟节点个数
	h.nodes = nodesArray{}

	for nodeKey, w := range h.weights { // 遍历每个真实节点 标识符和权重
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots))) // 节点权重/总权重 * 哈希环总虚拟节点个数 = 节点在哈希环上的虚拟节点个数
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i))) // 虚拟节点哈希值 = hash(真实节点标识符:虚拟节点下标i)
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   nodeKey,
				spotValue: genValue(hashBytes[6:10]), // TODO: 具体意义是什么？
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort()
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}

// 查询资源存储在哪一个虚拟节点上
func (h *HashRing) GetNode(s string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := genValue(hashBytes[6:10])                                                        // 资源的哈希值
	i := sort.Search(len(h.nodes), func(i int) bool { return h.nodes[i].spotValue >= v }) // 查找哈希环上哈希值 >= 资源哈希值 的第一个虚拟节点

	if i == len(h.nodes) { // 哈希环是一个环形结构, 如果资源哈希值 > 哈希环 0~len(h.nodes)-1 上所有虚拟节点哈希值，那么在顺时针方向上，存储该资源的就是第0个虚拟节点
		i = 0
	}
	return h.nodes[i].nodeKey // 返回查询到的虚拟节点对应的真实节点的标识符
}
