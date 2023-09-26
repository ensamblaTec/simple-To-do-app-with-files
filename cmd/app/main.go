package main

import (
	"ProyectoFinal/database"
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
		case "2":
		case "3":
		case "4":
			return
		default:
			println("invalid input... try again.")
		}
	}
}
