package repo

import (
	"fmt"
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) GetTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tempTasks := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tempTasks[k] = v
	}

	return tempTasks
}

func (l *List) GetUnCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	ucTasks := make(map[string]Task)

	for k, v := range l.tasks {
		fmt.Println(k, v, v.IsCompleted)
		if !v.IsCompleted {
			ucTasks[k] = v
		}
	}

	return ucTasks
}

func (l *List) GetTaskByID(id string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (l *List) GetTasksIds() []string {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	ids := []string{}

	for key := range l.tasks {
		ids = append(ids, key)
	}

	return ids
}

func (l *List) AddTask(newTask Task) (string, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	for _, task := range l.tasks {
		if task.Title == newTask.Title {
			return "", ErrTaskAlreadyExists
		}
	}

	id := newTask.id
	l.tasks[id] = newTask
	return id, nil
}

func (l *List) DeleteTask(id string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, id)

	return nil
}

func (l *List) ChangeCompletedTask(id string, isCompleted bool) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[id]

	if !ok {
		return ErrTaskNotFound
	}

	task.ChangeCompleted(isCompleted)

	l.tasks[id] = task

	return nil
}
