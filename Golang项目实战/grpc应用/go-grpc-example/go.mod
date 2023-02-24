module github.com/go-grpc-example
go 1.19

replace (
 	github.com/go-grpc-example/proto/search => C:/Users/hp-pc/GolandProjects/go-grpc-example/proto/search
	github.com/go-grpc-example/proto/stream => C:/Users/hp-pc/GolandProjects/go-grpc-example/proto/stream
)

require (
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
