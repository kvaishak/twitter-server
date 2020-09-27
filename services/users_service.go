package services

import (
	"encoding/json"

	"github.com/kvaishak/twitter-server/dao"
	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsers() (*[]model.User, *errors.AppError) {
	usersArr, err := dao.GetUsers()
	if err != nil {
		return nil, err
	}

	return usersArr, nil
}

func GetUserData(username string) (*model.User, *errors.AppError) {
	userData, err := dao.GetUserData(username)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func CreateUser(reqBody []byte) (bool, *errors.AppError) {
	var newUser = model.NewUser{}
	json.Unmarshal(reqBody, &newUser)

	isCreated, err := dao.CreateUser(newUser)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}
