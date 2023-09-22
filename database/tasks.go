package database

import (
	"ProyectoFinal/internal/models"
	"bufio"
	"fmt"
	"strings"
)

func GetTask(id int) *models.Task {
	scanner := bufio.NewScanner(Tasks.file)
	for scanner.Scan() {
		linea := scanner.Text()
		separateLine := strings.Split(linea, "|")
		fmt.Println(separateLine)
	}
	return &models.Task{}
}

func filterID(info ...string) {

}
