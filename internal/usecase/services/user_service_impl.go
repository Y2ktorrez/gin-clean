package services

import (
	"gin-clean/internal/entity"
	"gin-clean/internal/usecase/repositories"
)

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(user *entity.User) error {
	return s.repo.Create(user)
}

func (s *UserServiceImpl) GetUserByID(id uint) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserServiceImpl) GetAllUsers() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *UserServiceImpl) UpdateUser(user *entity.User) error {
	return s.repo.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
