package model

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

type NewUser struct {
	UserId   string
	UserName string
}
