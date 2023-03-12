package geecache

// 一种只读的数据类型，作为缓存的数据
type ByteView struct {
	b []byte // 数据选择 byte 类型是为了能够支持任意的数据类型的存储，例如字符串、图片等。
}

// 实现Value接口必须实现Len()方法
func (bv ByteView) Len() int {
	return len(bv.b)
}

// 返回缓存数据的副本，防止被修改
func (bv ByteView) ByteSlice() []byte {
	return cloneBytes(bv.b)
}

// 完成缓存数据的拷贝,返回其副本
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// 可选性，以字符串方式获取缓存数据
func (bv ByteView) String() string {
	return string(bv.b)
}
