package domain

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Password string `json:"password"`
	Role     UserRoles
	Status   UserStatus
}

type UserStatus int32

const (
	UserStatus_ACTIVE UserStatus = iota
	UserStatus_INACTIVE
	UserStatus_BANNED
	UserStatus_SUSPENDED
)

type UserRoles int32

const (
	UserRoles_ADMIN UserRoles = iota
	UserRoles_USER
)
