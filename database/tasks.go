package database

import (
	"ProyectoFinal/internal/models"
	"bufio"
	"strconv"
	"strings"
)

func GetTask(id int) (*models.Task, error) {
	scanner := bufio.NewScanner(Tasks.file)
	for scanner.Scan() {
		line := scanner.Text()
		currentTask := strings.Split(line, "|")
		if currentTask[0] == strconv.Itoa(id) {
			task, err := models.CreateTask(currentTask...)
			if err != nil {
				return nil, err
			}
			return task, nil
		}
	}
	return &models.Task{}, nil
}

func GetTaskByUser(id int) (tasks []*models.Task, err error) {
	scanner := bufio.NewScanner(Tasks.file)
	for scanner.Scan() {
		line := scanner.Text()
		currentTask := strings.Split(line, "|")
		if currentTask[2] == strconv.Itoa(id) {
			task, err := models.CreateTask(currentTask...)
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}
	}
	return
}
