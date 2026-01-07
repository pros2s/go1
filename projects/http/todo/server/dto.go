package server

import (
	"encoding/json"
	"errors"
	"time"
)

// task

type CompleteTaskDTO struct {
	Completed bool `json:"completed"`
}

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDTO) Validate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

// error

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(mes string) ErrorDTO {
	return ErrorDTO{
		Message: mes,
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}
