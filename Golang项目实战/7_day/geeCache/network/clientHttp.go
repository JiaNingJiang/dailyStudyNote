package network

import (
	"fmt"
	pb "geeCache/geecachepb"
	"geeCache/peer"
	"io/ioutil"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/proto"
)

// 实现 PeerGetter 接口(客户端)
type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(in.Group),
		url.QueryEscape(in.Key),
	)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	// 将http响应报文中的protobuff数据UnMarshal,输入到out中
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return nil
}

var _ peer.PeerGetter = (*httpGetter)(nil)
