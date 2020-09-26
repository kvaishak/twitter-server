package dao

import (
	"fmt"
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsers() (*[]model.User, *errors.AppError) {
	fmt.Println("Go MySQL connectivity")
	db := DbConn()

	results, err := db.Query("select UserId, UserName, UserEmail, FirstName, LastName from usertbl;")
	defer db.Close()

	if results != nil {
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

	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Error in getting Users data from the database"),
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}
