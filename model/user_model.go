package model

type User struct {
	UserID    int64  `json:"userid"`
	UserName  string `json:"username"`
	Email     string `json:"useremail"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type NewUser struct {
	Username  string
	Useremail string
	Firstname string
	Lastname  string
}
