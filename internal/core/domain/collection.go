package domain

type Collection struct {
	Id       string `json:"Id"`
	Name     string
	UserId   string
	User     *User
	Services []*Service
}
