package user_service

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
)

type userService struct {
	repository ports.UserRepository
}

func (u userService) NewApp() error {
	//TODO implement me
	panic("implement me")
}

func (u userService) List() ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repository ports.UserRepository) ports.UserService {
	return &userService{repository: repository}
}
