package configs

import (
	"os"
	"strconv"
)

var DatabasePostgreSQL = databaseConfig{
	Host:     os.Getenv("Host"),
	Port:     getPort(),
	User:     os.Getenv("User"),
	Password: os.Getenv("Password"),
	DBName:   os.Getenv("DBName"),
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		return 5432
	}
	return port
}
