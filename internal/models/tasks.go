package models

import "time"

type Task struct {
	id                              uint
	task                            string
	idUser                          uint
	completed                       bool
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

func (task *Task) ChangeTitle(title string) bool {
	if len(title) == 0 {
		return false
	}
	task.task = title
	return true
}

func (task *Task) ChangeCompleted() {
	task.completed = !task.completed
}
