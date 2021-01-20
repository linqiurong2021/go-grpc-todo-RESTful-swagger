package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/linqiurong2021/todo/api/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client")
	//
	address := "127.0.0.1:8088"
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect %s error : %s\n", address, err)
	}
	defer con.Close()
	//
	client := v1.NewTodoServiceClient(con)
	resp, err := client.Create(context.Background(), &v1.CreateRequest{
		Todo: &v1.Todo{
			Title: "Hello Client",
			Note:  "Test",
		},
	})
	if err != nil {
		log.Printf(" request failure : error %s\n", err)
	}
	// 请求返回结果
	fmt.Printf("%#v\n", resp)
}
