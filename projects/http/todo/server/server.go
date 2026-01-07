package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var PORT = ":9000"

type TODOServer struct {
	todoHandlers *HTTPHandlers
}

func NewTodoServer(handlers *HTTPHandlers) *TODOServer {
	return &TODOServer{
		todoHandlers: handlers,
	}
}

func (s *TODOServer) StartTodoServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.todoHandlers.HandleCreateTask)
	router.Path("/tasks/{id}").Methods("GET").HandlerFunc(s.todoHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.todoHandlers.HandleGetAllUncompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.todoHandlers.HandleGetAllTasks)
	router.Path("/tasks/{id}").Methods("PATCH").HandlerFunc(s.todoHandlers.HandleCompleteTask)
	router.Path("/tasks/{id}").Methods("DELETE").HandlerFunc(s.todoHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(PORT, router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
