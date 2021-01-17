package services

import (
	"encoding/json"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kvaishak/twitter-server/dao"
	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

var mySigningKey = []byte("mysupersecretsigningkey")

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

func FollowUser(uid string, followeName string) (bool, *errors.AppError) {
	isFollowed, err := dao.FollowUser(uid, followeName)
	if err != nil {
		return false, err
	}

	return isFollowed, nil
}

func UnFollowUser(uid string, followeName string) (bool, *errors.AppError) {
	isFollowed, err := dao.UnFollowUser(uid, followeName)
	if err != nil {
		return false, err
	}

	return isFollowed, nil
}

func IsFollowing(uid string, followeName string) (bool, *errors.AppError) {
	isFollowed, err := dao.IsFollowing(uid, followeName)
	if err != nil {
		return false, err
	}

	return isFollowed, nil
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

func GenerateJWT(uid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)

	claims["uid"] = uid
	token.Claims = claims

	tokenStr, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Error in generation JWT : %s", err)
		return "", err
	}
	return tokenStr, nil
}

func GetUserJWT(uid string) (string, error) {
	tokenString, err := GenerateJWT(uid)

	if err != nil {
		fmt.Errorf("Error in generation JWT in the main funciton : %s", err.Error())
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}
