package todo

import (
	"fmt"

	"go1/projects/http/todo/repo"
	"go1/projects/http/todo/server"
)

func TestTodo() {
	todoList := repo.NewList()
	todoHTTPHandlers := server.NewHTTPHandlers(todoList)
	todoServer := server.NewTodoServer(todoHTTPHandlers)

	fmt.Println("Todo server is working...")
	if err := todoServer.StartTodoServer(); err != nil {
		fmt.Println("Error with starting server: ", err.Error())
	}
}
