package gee

import (
	"fmt"
	"testing"
)

func TestParsePattern(t *testing.T) {
	fmt.Println(parsePattern("/home/ShanDong/QingDao"))
}

func newTestRouter() *router {
	r := newRouter()
	//r.addRouter("GET", "/", nil)
	// r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/*/ccc", nil)
	// r.addRouter("GET", "/home/ShanDong/QingDao", nil)
	// r.addRouter("GET", "/hi/:name", nil)
	// r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/shandong/QingDao")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	// if n.pattern != "/home/ShanDong/QingDao" {
	// 	t.Fatal("should match /home/ShanDong/QingDao")
	// }

	fmt.Println(n.pattern)
	_ = ps
	// if ps["name"] != "geektutu" {
	// 	t.Fatal("name should be equal to 'geektutu'")
	// }

	// fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}
