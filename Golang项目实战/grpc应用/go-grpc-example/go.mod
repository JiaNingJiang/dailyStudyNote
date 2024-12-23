module github.com/go-grpc-example

go 1.19

replace (
	github.com/go-grpc-example/pkg/gtls => C:/Users/hp-pc/GolandProjects/go-grpc-example/pkg/gtls
	github.com/go-grpc-example/proto/search => C:/Users/hp-pc/GolandProjects/go-grpc-example/proto/search
	github.com/go-grpc-example/proto/stream => C:/Users/hp-pc/GolandProjects/go-grpc-example/proto/stream
)

require (
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
