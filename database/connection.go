package database

import (
	"fmt"
	"io"
	"os"
)

const TASKS_HEADERS = "|id|task|id_user|completed|created_at|updated_at|deleted_at|"
const USERS_HEADERS = "|id|email|password|created_at|updated_at|deleted_at|"

var users *File
var tasks *File

type File struct {
	name string
	file *os.File
}

type Files interface {
	CreateFile(string) (*os.File, error)
	OpenFile(string) error
	ReadFile() (string, error)
	WriteFile(string) error
	CloseFile()
}

func Init() error {
	create, err := CreateFile("tasks")
	if err != nil {
		return err
	}
	tasks = &File{
		name: "tasks",
		file: create,
	}

	create, err = os.Create("database/users.txt")
	if err != nil {
		return err
	}

	_, err = create.WriteString(USERS_HEADERS)
	if err != nil {
		return err
	}

	users = &File{
		name: "users",
		file: create,
	}

	return nil
}

func (file *File) CreateFile(fileName string) (*os.File, error) {
	create, err := os.Create("database/" + fileName)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (file *File) OpenFile(fileName string) error {
	var err error
	conn, err = os.Open("database/" + file)
	if err != nil {
		return err
	}
	return nil
}

func (file *File) ReadFile(conn Files) (string, error) {
	file, err := io.ReadAll(conn)
	if err != nil {
		return "", err
	}
	newString := string(file)
	return newString, nil
}

func (file *File) WriteFile(info string, database Files) error {
	rowsAffected, err := database.file.Write([]byte(info))
	if err != nil {
		return err
	}
	fmt.Printf("RowsAffected: %d", rowsAffected)
	return nil
}

func (file *File) CloseFile() {
	err := conn.Close()
	if err != nil {
		return
	}
}
