package main

import (
	"ProyectoFinal/database"
	"ProyectoFinal/internal/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	repeatSymbol = strings.Repeat("=", 10)
	startMenu    = (repeatSymbol + "MENU" + repeatSymbol + "\n") +
		"[1] Mostrar tareas\n" +
		"[2] Crear tarea\n" +
		"[3] Seleccionar tarea\n" +
		"[4] Salir\n" +
		">>> "
)

func main() {
	// init database files
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	// closes database
	defer func(Tasks *database.File) {
		err := Tasks.CloseFile()
		if err != nil {
			log.Fatal(err)
		}
	}(database.Users)
	defer func(Users *database.File) {
		err := Users.CloseFile()
		if err != nil {
			log.Fatal(err)
		}
	}(database.Tasks)
	Init()
}

func Init() {
	menuInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(startMenu)
		err := menuInput.Scan()
		if !err {
			continue
		}
		switch menuInput.Text() {
		case "1":
			info, err := database.GetTaskByUser(1)
			if err != nil {
				log.Fatalf("cannot see tasks: %s", err)
			}
			for index, value := range info {
				fmt.Println(index, value.GetTitle())
			}
		case "2":
			var task *models.Task
			task = CreateTask()
			if task == nil {
				log.Fatalf("cannot create task")
			}
			err := database.Tasks.AppendFile(task.ToString())
			if err != nil {
				log.Fatalf("cannot append file: %s", err)
			}
		case "3":
		case "4":
			return
		default:
			println("invalid input... try again.")
		}
	}
}

func CreateTask() (task *models.Task) {
	var title string
	var onBool bool = false

	for !onBool {
		title, onBool = Input()
	}

	return models.NewTask(title)
}

func Input() (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Title: ")
	err := scanner.Scan()
	if !err {
		return "", false
	}
	input := scanner.Text()
	if len(input) == 0 {
		return "", false
	}
	return input, true
}
