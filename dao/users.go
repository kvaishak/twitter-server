package dao

import (
	"fmt"

	"github.com/kvaishak/twitter-server/model"
)

func GetUsers() (*[]model.User, error) {
	fmt.Println("Go MySQL connectivity")
	db := DbConn()

	results, err := db.Query("select UserId, UserName, UserEmail, FirstName, LastName from usertbl;")
	if err != nil {
		panic(err.Error())
	}

	//Getting all users data
	usersArr := []model.User{}
	for results.Next() {
		var user model.User
		err = results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err.Error())
		}
		usersArr = append(usersArr, user)
	}

	defer db.Close()
	return &usersArr, nil
}
