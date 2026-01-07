package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	todo "go1/projects/http/todo/repo"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	// decode
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := NewErrorDTO(err.Error())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	// validate
	if err := taskDTO.Validate(); err != nil {
		errDTO := NewErrorDTO(err.Error())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	// create
	newTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	newTaskID, err := h.todoList.AddTask(newTask)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	taskToWrite, err := h.todoList.GetTaskByID(newTaskID)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(taskToWrite, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error with write new task: ", err.Error())
	}
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := h.todoList.GetTaskByID(id)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error with write response about getting task: ", err.Error())
	}
}

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.GetTasks()
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error with write all tasks: ", err.Error())
	}
}

func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	ucTasks := h.todoList.GetUnCompletedTasks()
	b, err := json.MarshalIndent(ucTasks, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error with write uncompleted tasks: ", err.Error())
	}
}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completedDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completedDTO); err != nil {
		errDTO := NewErrorDTO(err.Error())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	if err := h.todoList.ChangeCompletedTask(id, completedDTO.Completed); err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.Message, http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.Message, http.StatusInternalServerError)
		return
	}

	taskToWrite, err := h.todoList.GetTaskByID(id)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(taskToWrite, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error with write completed task: ", err.Error())
	}
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.todoList.DeleteTask(id); err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
