package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ugurcsen/sand-panel/configs"
)

type db struct {
	*sql.DB
}

func Connect() (*db, error) {
	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.DatabasePostgreSQL.Host,
		configs.DatabasePostgreSQL.Port,
		configs.DatabasePostgreSQL.User,
		configs.DatabasePostgreSQL.Password,
		configs.DatabasePostgreSQL.DBName))
	return &db{DB}, err
}
