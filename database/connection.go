package database

import (
	"fmt"
	"os"
)

const (
	FilesRoute   = "database/"
	TasksHeaders = "id|task|id_user|completed|created_at|updated_at|deleted_at"
	UsersHeaders = "id|email|password|created_at|updated_at|deleted_at"
)

var (
	Users *File
	Tasks *File
)

type File struct {
	name string
	file *os.File
}

func Init() error {
	name := "tasks.txt"
	Tasks = &File{
		name: name,
	}
	create, err := CreateFile(name)
	if err != nil {
		return err
	}
	_, err = create.WriteString(TasksHeaders)
	if err != nil {
		return err
	}
	err = create.Close()
	if err != nil {
		return err
	}
	Tasks.file, err = OpenFile(name)
	if err != nil {
		return err
	}

	name = "users.txt"
	Users = &File{
		name: name,
	}
	create, err = CreateFile(name)
	if err != nil {
		return err
	}
	_, err = create.WriteString(UsersHeaders)
	if err != nil {
		return err
	}
	err = create.Close()
	if err != nil {
		return err
	}
	Users.file, err = OpenFile(name)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(fileName string) (*os.File, error) {
	create, err := os.Create(FilesRoute + fileName)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func OpenFile(fileName string) (*os.File, error) {
	conn, err := os.OpenFile(FilesRoute+fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (file *File) ReadFile() (string, error) {
	info, err := os.ReadFile(FilesRoute + file.name)
	if err != nil {
		return "", err
	}
	newString := string(info)
	return newString, nil
}

func (file *File) AppendFile(info string) error {
	rowsAffected, err := file.file.WriteString("\n" + info)
	if err != nil {
		return err
	}
	fmt.Printf("RowsAffected: %d\n", rowsAffected)
	return nil
}

func (file *File) CloseFile() error {
	err := file.file.Close()
	if err != nil {
		return err
	}
	file.file = nil
	return nil
}
