package search

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *SearchRequest) (*SearchResponse, error) {
	for i := 0; i < 5; i++ {
		if ctx.Err() == context.Canceled { // 检查context是否已经超时销毁
			return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
		}

		time.Sleep(1 * time.Second) // 模拟超时
	}
	return &SearchResponse{Response: r.GetRequest() + " Server"}, nil

}

func (s *SearchService) mustEmbedUnimplementedSearchServiceServer() {
}

func (s *SearchService) GetResult(reqID uint32) uint32 {
	return (reqID + 100)
}
