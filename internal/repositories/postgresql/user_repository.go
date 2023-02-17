package postgresql

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"log"
)

var _ ports.UserRepository = &userRepository{}

var testData = []domain.User{
	{
		Id:       "1",
		Email:    "admin@example.com",
		Name:     "Admin",
		Surname:  "",
		Password: "1234",
		Role:     domain.UserRolesAdmin,
		Status:   domain.UserStatusActive,
	},
	{
		Id:       "2",
		Email:    "user@example.com",
		Name:     "User",
		Surname:  "",
		Password: "4321",
		Role:     domain.UserRolesUser,
		Status:   domain.UserStatusActive,
	},
}

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

func (p userRepository) List() ([]domain.User, error) {
	return testData, nil
	rows, err := p.db.Query("SELECT id, user_name, password, email FROM users")
	if err != nil {
		return nil, err
	}
	var users []domain.User
	tempUser := domain.User{}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&tempUser.Id, &tempUser.Email, &tempUser.Password, &tempUser.Email); err != nil {
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
