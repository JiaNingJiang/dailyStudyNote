package hashTable

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"math"
	"math/big"
)

// 根据样本量n和容许的错误率p计算布隆过滤器大小(bit位数)m
func CalBloomSize(sampleSize uint64, errRate float64) uint64 {
	var bloomBitsSize = float64(sampleSize) * math.Log(errRate) / (math.Log(2) * math.Log(2)) * (-1)
	return uint64(math.Ceil(bloomBitsSize)) // 小数要向上取整
}

// 根据布隆过滤器大小m和样本量n计算布隆过滤器需要的哈希函数的数量
func CalHashFuncNum(sampleSize, bloomSize uint64) uint64 {
	var k = math.Log(2) * float64(bloomSize) / float64(sampleSize)
	return uint64(math.Ceil(k))
}

// 根据实际的样本量n、布隆过滤器大小m、哈希函数数量k计算真正的错误率p
func CalErrRate(sampleSize, bloomSize, hashFuncNum uint64) float64 {
	var y = float64(sampleSize) * float64(hashFuncNum) / float64(bloomSize)
	return math.Pow(float64(1)-math.Pow(math.E, y*float64(-1)), float64(hashFuncNum))
}

// 同一个key值输入不同的seed，可以得到不一样的哈希值。模拟不同的哈希函数
func HMACWithSHA128(seed []byte, key []byte) []byte {
	//hmac512 := hmac.New(sha512.New, key)
	hmac512 := hmac.New(sha1.New, seed)
	hmac512.Write(key)
	return hmac512.Sum(nil)
}

type BloomFilter struct {
	SampleSize  uint64  // 样本量
	BloomSize   uint64  // 大小(bit位数)
	HashFuncNum uint64  // 哈希函数数量
	ErrRate     float64 // 容许的错误率

	bitMap *BitMap         // 实现布隆过滤器的位图
	keys   map[uint32]bool // keys用来存储k个哈希函数需要用到的随机数
}

func NewBloomFilter(sampleSize uint64, errRate float64) *BloomFilter {
	bf := new(BloomFilter)
	bf.SampleSize = sampleSize
	bf.BloomSize = CalBloomSize(sampleSize, errRate)
	bf.HashFuncNum = CalHashFuncNum(sampleSize, bf.BloomSize)
	bf.ErrRate = CalErrRate(sampleSize, bf.BloomSize, bf.HashFuncNum)

	bf.bitMap = NewBitMap(bf.BloomSize)

	//是否是类似HMAC-SHA256那种通过改变seed值形成不同的哈希函数
	bf.keys = make(map[uint32]bool)
	for uint64(len(bf.keys)) < bf.HashFuncNum {
		randNum, err := rand.Int(rand.Reader, new(big.Int).SetUint64(math.MaxUint32))
		if err != nil {
			panic(err)
		}
		bf.keys[uint32(randNum.Uint64())] = true
	}
	return bf
}

// 向布隆过滤器中添加一条数据
func (bf *BloomFilter) Add(elem []byte) {
	var buf [4]byte
	for k := range bf.keys { // 用k个哈希函数计算该数据
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(buf[:], elem))     // 得到数据的哈希值
		result := hashResult.Mod(hashResult, big.NewInt(int64(bf.BloomSize))) // 哈希值 % m
		//把result对应的bit置1
		if err := bf.bitMap.SetBitOne(result.Uint64()); err != nil {
			panic(err)
		}
	}
}

// 判断数据是否在布隆过滤器里面
func (bf *BloomFilter) IsContain(elem []byte) bool {
	var buf [4]byte
	for k := range bf.keys {
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(buf[:], elem))     // 计算哈希值
		result := hashResult.Mod(hashResult, big.NewInt(int64(bf.BloomSize))) // 哈希值 % m
		//查询result对应的bit是否被置1
		if status, err := bf.bitMap.CheckBitStatus(result.Uint64()); err != nil {
			panic(err)
		} else {
			if !status { // 只要有一个哈希函数的结果不在布隆过滤器的位图中，那么该数据就不存在于布隆过滤器中
				return false
			}
		}
	}
	return true
}
