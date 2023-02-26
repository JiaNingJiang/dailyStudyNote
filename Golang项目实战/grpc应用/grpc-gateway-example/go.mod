module github.com/grpc-gateway-example

go 1.19

replace (
	github.com/grpc-gateway-example/client => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/client
	github.com/grpc-gateway-example/cmd => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/cmd
	github.com/grpc-gateway-example/pkg => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/pkg
	github.com/grpc-gateway-example/proto => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/proto
	github.com/grpc-gateway-example/server => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/server

	github.com/grpc-gateway-example/pkg/utils => C:/Users/hp-pc/GolandProjects/grpc-gateway-example/pkg/utils
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.1
	github.com/spf13/cobra v1.6.1
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
)
