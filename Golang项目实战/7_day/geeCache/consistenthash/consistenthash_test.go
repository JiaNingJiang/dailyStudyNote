package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	// 要进行测试，那么我们需要明确地知道每一个传入的 key 的哈希值，那使用默认的 crc32.ChecksumIEEE 算法显然达不到目的。所以在这里使用了自定义的 Hash 算法。
	// 自定义的 Hash 算法只处理数字，传入字符串表示的数字，返回对应的数字即可。
	hash := NewHashRing(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2") // 构建出的哈希环应该是： 2, 4, 6, 12, 14, 16, 22, 24, 26

	testCases := map[string]string{ // key为缓存资源，value为真实节点的名称
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2", // 资源27的hash值将比任何虚拟节点的hash值都大，因此会存储在m.keys[0]
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// Adds 8, 18, 28
	hash.Add("8")

	// 27 should now map to 8.
	testCases["27"] = "8" // 资源27将会被重新缓存到真实节点8

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

}
