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
	UserStatusActive UserStatus = iota
	UserStatusInactive
	UserStatusBanned
	UserStatusSuspended
)

func (s UserStatus) String() string {
	return [...]string{"ACTIVE", "INACTIVE", "BANNED", "SUSPENDED"}[s]
}

type UserRoles int32

const (
	UserRolesAdmin UserRoles = iota
	UserRolesUser
)
