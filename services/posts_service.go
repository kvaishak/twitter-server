package services

import (
	"encoding/json"

	"github.com/kvaishak/twitter-server/dao"
	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsersPost(username string) (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetUsersPost(username)
	if err != nil {
		return nil, err
	}

	return postsData, nil
}

func GetAllUsersPost() (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetAllUsersPost()
	if err != nil {
		return nil, err
	}

	return postsData, nil
}

func GetFollowersPost(uid string) (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetFollowersPost(uid)
	if err != nil {
		return nil, err
	}

	return postsData, nil
}

func GetFollowersTimedPost(uid string, lastTweetId string) (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetFollowersTimedPost(uid, lastTweetId)
	if err != nil {
		return nil, err
	}

	return postsData, nil
}

func NewPost(reqBody []byte) (bool, *errors.AppError) {
	var newPostData = model.NewPost{}
	json.Unmarshal(reqBody, &newPostData)

	isCreated, err := dao.NewPost(newPostData)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}
