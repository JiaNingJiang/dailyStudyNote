package search

import (
	"context"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *SearchRequest) (*SearchResponse, error) {
	return &SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func (s *SearchService) mustEmbedUnimplementedSearchServiceServer() {
}

func (s *SearchService) GetResult(reqID uint32) uint32 {
	return (reqID + 100)
}
