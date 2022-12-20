package user_service

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
)

type userService struct {
	repository ports.UserRepository
}

func (u userService) List() ([]*domain.User, error) {
	return u.repository.List()
}

func NewUserService(repository ports.UserRepository) ports.UserService {
	return &userService{repository: repository}
}
