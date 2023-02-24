package stream

import (
	"io"
	"log"
)

type StreamService struct {
}

func (s *StreamService) List(r *StreamRequest, stream StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&StreamResponse{
			Pt: &StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *StreamService) Record(stream StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&StreamResponse{Pt: &StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}
func (s *StreamService) Route(stream StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&StreamResponse{ // 发送一次
			Pt: &StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv() // 接收一次
		if err == io.EOF {      // 会因为客户端的退出而终止
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}
func (s *StreamService) mustEmbedUnimplementedStreamServiceServer() {

}
