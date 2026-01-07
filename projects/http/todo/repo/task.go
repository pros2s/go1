package repo

import (
	"time"

	"github.com/lithammer/shortuuid"
)

type Task struct {
	id          string
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`

	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

func NewTask(title string, description string) Task {
	return Task{
		id:          shortuuid.New(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) ChangeCompleted(isCompleted bool) {
	t.IsCompleted = isCompleted

	if isCompleted {
		completeTime := time.Now()
		t.CompletedAt = &completeTime
	} else {
		t.CompletedAt = nil
	}
}
