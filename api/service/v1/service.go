package v1

import (
	"context"
	"database/sql"
	"fmt"

	v1 "github.com/linqiurong2021/todo/api/proto/v1"
)

// TodoService TodoService
type TodoService struct {
	db *sql.DB
}

// Todo  Todo
type Todo struct {
	ID    int64
	Title string
	Note  string
}

// NewTodoService NewTodoService
func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

// Create Create
func (t *TodoService) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// 数据库连接
	rows, err := t.db.QueryContext(ctx, "select id,title,note from todo where id=? limit 1", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//
	todo := new(Todo)
	if rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Note) // 取决于select 的字段
		if err != nil {
			fmt.Printf("query error %s\n", err)
			return nil, err
		}
	}
	fmt.Printf("get Result %#v\n", todo)
	return &v1.CreateResponse{
		Code:    200,
		Message: "succcess",
		Todo: &v1.Todo{
			Title: todo.Title,
			Note:  todo.Note,
		},
	}, nil
}

// Update Update
func (t *TodoService) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// 数据库连接
	rows, err := t.db.QueryContext(ctx, "select id,title,note from todo where id=? limit 1", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//
	todo := new(Todo)
	if rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Note) // 取决于select 的字段
		if err != nil {
			fmt.Printf("query error %s\n", err)
			return nil, err
		}
	}
	fmt.Printf("get Result %#v\n", todo)
	return &v1.UpdateResponse{
		Code:    200,
		Message: "succcess",
		Todo: &v1.Todo{
			Title: todo.Title,
			Note:  todo.Note,
		},
	}, nil
}
