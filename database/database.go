package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ugurcsen/sand-panel/configs"
)

type DB struct {
	*sql.DB
}

var psqlInfo = ""

func connectionString() {
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.DatabasePostgreSQL.Host,
		configs.DatabasePostgreSQL.Port,
		configs.DatabasePostgreSQL.User,
		configs.DatabasePostgreSQL.Password,
		configs.DatabasePostgreSQL.DBName)
}

func Connect() (*DB, error) {
	if psqlInfo == "" {
		connectionString()
	}
	db, err := sql.Open("postgres", psqlInfo)
	return &DB{db}, err
}
