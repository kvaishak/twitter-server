package services

import (
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
