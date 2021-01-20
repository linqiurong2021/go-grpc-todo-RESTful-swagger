package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/linqiurong2021/todo/api/proto/v1"
	"github.com/linqiurong2021/todo/conf"
	"google.golang.org/grpc"
)

// Update

var (
	// command-line options:
	// gRPC server endpoint
	// 需要修改 按端口号 localhost:8088
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8088", "gRPC server endpoint")
)

// 开启swagger
func runSwagger(gwmux *runtime.ServeMux) error {
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	// 启用swagger-ui
	func(mux *http.ServeMux) {
		// 设置目录
		dir := "../../ui/swagger/3.40"
		mux.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir(dir))))
	}(mux)
	// swagger 端口号
	swaggerPort := conf.Conf.AppConfig.SwaggerPort
	return http.ListenAndServe(swaggerPort, mux)
}

// 启用代理
func runProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 需要修改 按服务名称
	err := gw.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	if err != nil {
		return err
	}
	// 启用swagger
	go runSwagger(mux)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	// swagger 端口号
	gRPCPort := conf.Conf.AppConfig.GRPCPort
	return http.ListenAndServe(gRPCPort, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := runProxy(); err != nil {
		glog.Fatal(err)
	}
}
