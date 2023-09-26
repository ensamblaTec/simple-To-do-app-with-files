package models

import (
	"ProyectoFinal/pkg/utils"
	"strconv"
	"time"
)

type Task struct {
	id                              uint
	task                            string
	idUser                          uint
	completed                       bool
	tasks                           []*Task
	createdAt, updatedAt, deletedAt time.Time
}

func CreateTask(info ...string) (*Task, error) {
	if len(info) == 2 {
		convert, err := strconv.Atoi(info[1])
		if err != nil {
			return nil, err
		}
		return &Task{
			task:      info[0],
			idUser:    uint(convert),
			completed: false,
			createdAt: time.Now(),
		}, nil
	}
	convertID, err := strconv.Atoi(info[0])
	if err != nil {
		return nil, err
	}
	convertIDUser, err := strconv.Atoi(info[2])
	if err != nil {
		return nil, err
	}
	convertBool, err := strconv.ParseBool(info[3])
	if err != nil {
		return nil, err
	}
	createdAt, err := utils.StringToDate(info[4])
	if err != nil {
		return nil, err
	}
	if len(info) > 5 {
		updatedAt, _ := utils.StringToDate(info[5])
		deletedAt, _ := utils.StringToDate(info[6])
		return &Task{
			id:        uint(convertID),
			task:      info[1],
			idUser:    uint(convertIDUser),
			completed: convertBool,
			createdAt: createdAt,
			updatedAt: updatedAt,
			deletedAt: deletedAt,
		}, nil
	}
	return &Task{
		id:        uint(convertID),
		task:      info[1],
		idUser:    uint(convertIDUser),
		completed: convertBool,
		createdAt: createdAt,
	}, nil
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
