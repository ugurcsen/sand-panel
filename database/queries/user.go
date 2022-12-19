package queries

// User is SQL queries for users table
var User = userQuery{
	Select:                      `SELECT id, user_name, password, email FROM users WHERE id=$1`,
	SelectByUserNameAndPassword: `SELECT id, user_name, password, email FROM users WHERE user_name=$1 AND password=$2`,
	Insert:                      `INSERT INTO users (user_name, password, email) VALUES ($1, $2, $3) RETURNING id, user_name, password, email`,
	Update:                      `UPDATE users SET user_name=$2, password=$3, email=$4 WHERE id=$1`,
	Delete:                      `DELETE FROM users WHERE id=$1`,
}

// userQuery is SQL query structure for users table
type userQuery struct {
	Select                      string
	SelectByUserNameAndPassword string
	Insert                      string
	Update                      string
	Delete                      string
}
