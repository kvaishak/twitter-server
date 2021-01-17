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

	results, err := db.Query("select userId, userName from usertbl;")
	defer db.Close()

	if results != nil {
		//Getting all users data
		usersArr := []model.User{}
		for results.Next() {
			var user model.User
			err = results.Scan(&user.UserId, &user.UserName)
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
	err := db.QueryRow("select userId, userName from usertbl where userName=?", username).Scan(&userData.UserId, &userData.UserName)
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

func FollowUser(uid string, followerName string) (bool, *errors.AppError) {
	db := DbConn()

	_, err := db.Exec("insert into followstbl (userId, followerId) values (?, (select userId from usertbl where userName=?))", uid, followerName)
	// defer db.Close()

	if err != nil {
		return false, &errors.AppError{
			Message:    fmt.Sprintf("Error in following the user: %s's data", followerName),
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	} else {
		return true, nil
	}
}

func UnFollowUser(uid string, followerName string) (bool, *errors.AppError) {
	db := DbConn()

	_, err := db.Exec("delete from followstbl where userId=? AND followerId=(select userId from usertbl where userName=?)", uid, followerName)
	// defer db.Close()

	if err != nil {
		return false, &errors.AppError{
			Message:    fmt.Sprintf("Error in unfollowing the user: %s's data", followerName),
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	} else {
		return true, nil
	}
}

func IsFollowing(uid string, followerName string) (bool, *errors.AppError) {
	db := DbConn()

	var followerId int
	err := db.QueryRow("select followId from followstbl where userId=? AND followerId=(select userId from usertbl where userName=?);", uid, followerName).Scan(&followerId)
	defer db.Close()

	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, &errors.AppError{
			Message:    fmt.Sprintf("Error in getting the user: %s's follower data", followerName),
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	default:
		return true, nil
	}

}

func CreateUser(newUser model.NewUser) (bool, *errors.AppError) {
	db := DbConn()

	_, err := db.Exec("INSERT INTO usertbl (userId, userName) VALUES (?, ?)", newUser.UserId, newUser.UserName)
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
