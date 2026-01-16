package sql

import "time"

type TaskModel struct {
	ID           int
	Title        string
	Description  string
	Created_at   time.Time
	Completed    bool
	Completed_at *time.Time
}

func NewTask(title string, desc string) TaskModel {
	return TaskModel{
		Title:        title,
		Description:  desc,
		Created_at:   time.Now(),
		Completed:    false,
		Completed_at: nil,
	}
}
