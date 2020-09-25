package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsers() (*[]model.User, error) {
	fmt.Println("Go MySQL connectivity")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/twitterdb2")
	if err != nil {
		panic(err.Error())
	}

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

	return &usersArr, nil
}
