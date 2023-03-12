package geecache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expectRes := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(expectRes, v) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}
var loadCounts = make(map[string]int, len(db))

func mockGetter(key string) ([]byte, error) {
	fmt.Println("[MockDB] search key", key)

	if v, ok := db[key]; ok { // 获取的数据必须要存在于数据源中
		if _, ok := loadCounts[key]; !ok { // 如果要获取的数据不存在历史查询记录,将其查询次数初始化为0
			loadCounts[key] = 0
		}
		loadCounts[key]++ // 查询次数+1
		return []byte(v), nil

	} else {
		return nil, fmt.Errorf("%s not exist", key)
	}
}

func TestGroupGet(t *testing.T) {
	gee := NewGroup("scores", 2<<10, GetterFunc(mockGetter))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v { // 第一次查询用于缓存
			t.Fatal("failed to get value of Tom")
		} // load from callback function
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 { // 第二次查询会命中缓存
			t.Fatalf("cache %s miss", k)
		} // cache hit
	}

	if view, err := gee.Get("unknown"); err == nil { // 查询不存在的数据，不会发生缓存
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	} else {
		t.Log(err)
	}
}
