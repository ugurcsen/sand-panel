package postgresql

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"log"
)

type userRepository struct {
	db *db
}

func NewUserRepository() (ports.UserRepository, error) {
	DB, err := Connect()
	return &userRepository{db: DB}, err
}

func (p userRepository) Get(id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p userRepository) GetByUserNameAndPassword(userName, password string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p userRepository) List() ([]*domain.User, error) {
	rows, err := p.db.Query("SELECT id, user_name, password, email FROM users")
	users := make([]*domain.User, 0)
	tempUser := &domain.User{}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tempUser.ID, &tempUser.UserName, &tempUser.Password, &tempUser.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, tempUser)
	}

	return users, nil
}

func (p userRepository) Create(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p userRepository) Update(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p userRepository) Delete(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
