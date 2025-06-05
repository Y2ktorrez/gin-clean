package services

import "gin-clean/internal/entity"

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}
