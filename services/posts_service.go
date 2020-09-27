package services

import (
	"github.com/kvaishak/twitter-server/dao"
	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetFollowersPost(username string) (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetFollowersPost(username)
	if err != nil {
		return nil, err
	}

	return postsData, nil
}

func GetFollowersTimedPost(username string, lastPublishTime string) (*[]model.Post, *errors.AppError) {
	postsData, err := dao.GetFollowersTimedPost(username, lastPublishTime)
	if err != nil {
		return nil, err
	}

	return postsData, nil
}
