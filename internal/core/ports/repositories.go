package ports

import "github.com/ugurcsen/sand-panel/internal/core/domain"

type UserRepository interface {
	Get(id int64) (*domain.User, error)
	GetByUserNameAndPassword(userName, password string) (*domain.User, error)
	List() ([]*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(user *domain.User) error
}
