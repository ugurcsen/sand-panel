package postgresql

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
)

type postgresqlRepository struct{}

func NewPostgresqlRepository() ports.UserRepository {
	return &postgresqlRepository{}
}

func (p postgresqlRepository) Get(id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresqlRepository) GetByUserNameAndPassword(userName, password string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresqlRepository) List() (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresqlRepository) Create(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresqlRepository) Update(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresqlRepository) Delete(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
