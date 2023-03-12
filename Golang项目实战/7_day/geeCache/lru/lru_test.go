package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int { // 数据要进行存储,必须实现Value接口(实现Len()方法)
	return len(s)
}

func TestGet(t *testing.T) {

	lru := NewCache(0, nil) // 创建一个空Cache,容量上限为0

	lru.Add("key1", String("1234"))

	if value, ok := lru.Get("key1"); !ok || string(value.(String)) != "1234" {
		t.Fatal("cache hit key1=1234 failed")
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}

}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"

	cap := len(k1 + k2 + v1 + v2)
	lru := NewCache(cap, nil)

	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatal("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)

	callBack := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := NewCache(10, callBack)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}

}
