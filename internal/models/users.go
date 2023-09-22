package models

import (
	"ProyectoFinal/pkg/utils"
	"errors"
	"time"
)

type User struct {
	id                              uint
	email, password                 string
	createdAt, updatedAt, deletedAt time.Time
}

func CreateUser(email, password string) (*User, error) {
	if len(email) == 0 {
		return nil, errors.New("verify your email")
	}
	if len(password) == 0 {
		return nil, errors.New("verify your password")
	}
	hashedPassword, err := utils.Hashing256(password)
	if err != nil {
		return nil, errors.New("verify your password")
	}
	return &User{
		id:        1,
		email:     email,
		password:  hashedPassword,
		createdAt: time.Now(),
	}, nil
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) UpdateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("verify your email")
	}
	user.email = email
	user.updatedAt = time.Now()
	return nil
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) UpdatePassword(password string) error {
	if len(password) == 0 {
		return errors.New("verify your password")
	}
	hashedPassword, err := utils.Hashing256(password)
	if err != nil {
		return err
	}
	user.password = hashedPassword
	user.updatedAt = time.Now()
	return nil
}
