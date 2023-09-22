package main

import (
	"ProyectoFinal/database"
	"fmt"
	"log"
)

func main() {
	err := database.Init("tasks.txt")
	if err != nil {
		log.Println(err)
	}
	defer database.CloseFile()
	file, err := database.ReadFile()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file)

	err = database.WriteFile("\nNuevos datos")
	if err != nil {
		log.Println(err)
	}

	file, err = database.ReadFile()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file)
}
