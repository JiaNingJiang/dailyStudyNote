package utils

import (
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

//func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
//	if otherHandler == nil {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			grpcServer.ServeHTTP(w, r)
//		})
//	}
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
//			grpcServer.ServeHTTP(w, r)
//		} else {
//			otherHandler.ServeHTTP(w, r)
//		}
//	})
//}

func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
