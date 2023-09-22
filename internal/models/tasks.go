package models

import "time"

type Task struct {
	id                              uint
	task                            string
	idUser                          uint
	completed                       bool
	tasks                           []*Task
	createdAt, updatedAt, deletedAt time.Time
}

func CreateTask(task string, idUser uint) *Task {
	return &Task{
		task:      task,
		idUser:    idUser,
		completed: false,
		createdAt: time.Now(),
	}
}

func (task *Task) GetTitle() string {
	return task.task
}

func (task *Task) UpdateTitle(title string) bool {
	if len(title) == 0 {
		return false
	}
	task.task = title
	return true
}

func (task *Task) Completed() {
	task.completed = !task.completed
}

func (task *Task) isCompleted() bool {
	return task.completed
}
