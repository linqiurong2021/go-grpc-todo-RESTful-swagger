package server

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	pbV1 "github.com/linqiurong2021/todo/api/proto/v1"
	serviceV1 "github.com/linqiurong2021/todo/api/service/v1"
	"github.com/linqiurong2021/todo/conf"
	"google.golang.org/grpc"
)

// TodoServer TodoServer
type TodoServer struct {
}

// 初始化数据库连接
func initDB() *sql.DB {
	params := "charset=utf8&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", conf.Conf.DBConfig.User, conf.Conf.DBConfig.Password, conf.Conf.DBConfig.Host, conf.Conf.DBConfig.Sechma, params)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("connect database error ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping error ", err)
	}
	return db
}

// NewTodoServer NewTodoServer
func NewTodoServer(port string) error {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	// 数据库初始化
	db := initDB()
	defer db.Close()
	//
	service := serviceV1.NewTodoService(db)
	server := grpc.NewServer()
	//
	pbV1.RegisterTodoServiceServer(server, service)
	// 启用服务
	err = server.Serve(listen)
	if err != nil {
		return err
	}
	log.Printf("server start at %s\n", port)
	return nil
}
