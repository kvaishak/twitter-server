package dao

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsers() (*[]model.User, *errors.AppError) {

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
		Message:    "Error in getting Users data from the database",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}

func GetUserData(username string) (*model.User, *errors.AppError) {
	db := DbConn()

	userData := model.User{}
	err := db.QueryRow("select UserId, UserName, UserEmail, FirstName, LastName from usertbl where UserName=?", username).Scan(&userData.UserID, &userData.UserName, &userData.Email, &userData.FirstName, &userData.LastName)
	defer db.Close()

	switch {
	case err == sql.ErrNoRows:
		return nil, &errors.AppError{
			Message:    fmt.Sprintf("No user with username - %s in the Database", username),
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	case err != nil:
		return nil, &errors.AppError{
			Message:    fmt.Sprintf("Error in getting the user: %s's data", username),
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	default:
		return &userData, nil
	}
}

func CreateUser(newUser model.NewUser) (bool, *errors.AppError) {
	db := DbConn()

	_, err := db.Exec("INSERT INTO usertbl (UserName, UserPassword, UserEmail, FirstName, LastName) VALUES (?, ?, ?, ?, ?)", newUser.Username, newUser.Username, newUser.Useremail, newUser.Firstname, newUser.Lastname)
	if err != nil {
		return false, &errors.AppError{
			Message:    "Error in creating new User",
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	} else {
		return true, nil
	}

}
