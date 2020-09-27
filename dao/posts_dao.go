package dao

import (
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetFollowersPost(username string) (*[]model.Post, *errors.AppError) {

	db := DbConn()

	results, err := db.Query("SELECT TweetId, TweetText, PubTime, UserName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.TweetAuthorID=usertbl.UserId WHERE TweetAuthorID IN (SELECT FollowerId FROM followstbl WHERE UserId IN (SELECT UserId FROM usertbl WHERE UserName=?)) ORDER BY PubTime DESC LIMIT 0, 3;", username)

	defer db.Close()

	if results != nil {
		//Getting all Followers Post data
		postsArr := []model.Post{}
		for results.Next() {
			var post model.Post
			err = results.Scan(&post.TweetID, &post.TweetText, &post.PublishTime, &post.UserName)
			if err != nil {
				panic(err.Error())
			}
			postsArr = append(postsArr, post)
		}

		return &postsArr, nil
	}

	return nil, &errors.AppError{
		Message:    "Error in getting Tweets data from the database",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}

func GetFollowersTimedPost(username string, lastPublishTime string) (*[]model.Post, *errors.AppError) {

	db := DbConn()
	results, err := db.Query("SELECT TweetId, TweetText, PubTime, UserName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.TweetAuthorID=usertbl.UserId WHERE TweetAuthorID IN (SELECT FollowerId FROM followstbl WHERE UserId IN (SELECT UserId FROM usertbl WHERE UserName=?)) AND PubTime<? ORDER BY PubTime DESC LIMIT 0, 3;", username, lastPublishTime)

	defer db.Close()

	if results != nil {
		//Getting all Followers Post data
		postsArr := []model.Post{}
		for results.Next() {
			var post model.Post
			err = results.Scan(&post.TweetID, &post.TweetText, &post.PublishTime, &post.UserName)
			if err != nil {
				panic(err.Error())
			}
			postsArr = append(postsArr, post)
		}

		return &postsArr, nil
	}

	return nil, &errors.AppError{
		Message:    "Error in getting Tweets data from the database",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}
