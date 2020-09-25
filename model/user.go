package model

type User struct {
	UserID    int64  `json:"userid"`
	UserName  string `json:"username"`
	Email     string `json:"useremail"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
