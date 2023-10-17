package services

import (
	"strings"

	"github.com/bitebait/cupcakestore/models"
	"github.com/bitebait/cupcakestore/repositories"
)

type UserService interface {
	Create(user *models.User) error
	FindAll(p *models.Pagination, filter string) []*models.User
	FindById(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create(user *models.User) error {
	s.normalizeUser(user)
	return s.userRepository.Create(user)
}

func (s *userService) FindAll(p *models.Pagination, filter string) []*models.User {
	return s.userRepository.FindAll(p, filter)
}

func (s *userService) FindById(id uint) (*models.User, error) {
	return s.userRepository.FindById(id)
}

func (s *userService) FindByUsername(username string) (*models.User, error) {
	return s.userRepository.FindByUsername(username)
}

func (s *userService) Update(user *models.User) error {
	s.normalizeUser(user)
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id uint) error {
	user, err := s.FindById(id)
	if err != nil {
		return err
	}
	return s.userRepository.Delete(user)
}

func (s *userService) normalizeUser(user *models.User) {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
}
