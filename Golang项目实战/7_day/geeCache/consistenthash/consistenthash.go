package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32 // 自定义Hash函数,求取数据的Hash值

type Map struct {
	hash     Hash           // Hash函数
	replicas int            // 每个真实节点虚拟化的节点个数
	keys     []int          //  哈希环，所有虚拟节点的Hash集合(完成sort排序)
	hashMap  map[int]string // 存储虚拟节点Hash与真实节点名称的映射关系
}

func NewHashRing(replicas int, hash Hash) *Map {
	m := &Map{
		hash:     hash,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if hash == nil {
		m.hash = crc32.ChecksumIEEE // 默认的Hash算法为crc32.ChecksumIEEE
	}

	return m
}

// 根据nodeID集合构建一致性哈希环
func (m *Map) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ { // 每个真实节点都需要创建replicas个虚拟节点
			hash := int(m.hash([]byte(strconv.Itoa(i) + node))) // 求每一个虚拟节点的Hash值(虚拟节点的名称是：strconv.Itoa(i) + node，即通过添加编号的方式区分同个真实节点的不同虚拟节点。)
			m.keys = append(m.keys, hash)                       //存入哈希环
			m.hashMap[hash] = node                              // 记录虚拟节点和真实节点的映射关系
		}
	}

	sort.Ints(m.keys) // 对哈希环进行排序
}

// 为待缓存资源进行进行节点选择
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key))) // 求缓存资源的Hash

	index := sort.Search(len(m.keys), func(i int) bool { // 获取符合节点在m.keys中的下标。查询的终止条件是index == len(m.keys),此时意味着资源的Hash比所有虚拟节点的Hash值都要大
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[index%len(m.keys)]] // 求余的目的是为了当 index == len(m.keys) 时，将资源存储在m.keys[0]
}
