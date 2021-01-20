# 安装 参考  https://github.com/grpc-ecosystem/grpc-gateway

参考: https://blog.csdn.net/wangjunsheng/article/details/80779276

```go
go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u -v github.com/golang/protobuf/protoc-gen-go
```




# 找不到import 
```
解决方案: 
1: 打开 https://github.com/protocolbuffers/protobuf 找到相应的版本下载并解压
2: 把解压后的文件夹中的google复制到 /usr/local/include/
或
1: 把需要的文件放到当前目录下
2: 指定目录为当前目录下

```




# 生成pd文件

protoc -I /usr/local/include -I .  --go_out=plugins=grpc:. todo-service.proto

# 生成gw文件
protoc -I/usr/local/include -I. -I $GOPATH/bin --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway  --grpc-gateway_out=logtostderr=true:. *.proto


# 同时生成pd 与 gw.pd 与 grpc.pd

protoc -I /usr/local/include -I .  --go_out=. --plugin=protoc-gen-go-grpc=$GOPATH/bin/protoc-gen-go-grpc --plugin=protoc-gen-go-grpc=$GOBIN/protoc-gen-go-grpc --go-grpc_out=.  *.proto 

protoc -I /usr/local/include -I .  --go_out=. --plugin=protoc-gen-go-grpc=$GOBIN/protoc-gen-go-grpc --plugin=protoc-gen-go-grpc=$GOBIN/protoc-gen-go-grpc --go-grpc_out=.  *.proto 

# 生成openapi
当前目录下生成
protoc --plugin=protoc-gen-swagger=$GOBIN/protoc-gen-swagger --swagger_out=logtostderr=true:. *.proto

生成到指定目录
protoc --plugin=protoc-gen-swagger=$GOBIN/protoc-gen-swagger --swagger_out=logtostderr=true:../../../ui/swagger/3.40/ *.proto