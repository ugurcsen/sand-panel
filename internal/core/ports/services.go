package ports

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
)

type UserService interface {
	List() ([]domain.User, error)
}
