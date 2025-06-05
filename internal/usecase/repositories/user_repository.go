package repositories

import "gin-clean/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id uint) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}
