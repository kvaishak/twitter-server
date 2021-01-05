package dao

import (
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetUsersPost(username string) (*[]model.Post, *errors.AppError) {

	db := DbConn()

	results, err := db.Query("SELECT tweetId, tweetText, pubTime, userName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.tweetAuthorId=usertbl.userId WHERE tweetAuthorId IN (SELECT userId FROM usertbl WHERE userName=?) ORDER BY TweetId DESC;", username)

	defer db.Close()

	if results != nil {
		//Getting all Users Post data
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

func GetAllUsersPost() (*[]model.Post, *errors.AppError) {

	db := DbConn()

	results, err := db.Query("SELECT tweetId, tweetText, pubTime, userName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.tweetAuthorId=usertbl.userId ORDER BY tweetId DESC;")

	defer db.Close()

	if results != nil {
		//Getting all Users Post data
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
		Message:    "Error in getting All Tweets data from the database",
		StatusCode: http.StatusNotFound,
		Status:     "not found",
	}

}

func GetFollowersPost(username string) (*[]model.Post, *errors.AppError) {

	db := DbConn()

	results, err := db.Query("SELECT tweetId, tweetText, pubTime, userName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.tweetAuthorId=usertbl.userId WHERE tweetAuthorId IN (SELECT followerId FROM followstbl WHERE userId IN (SELECT userId FROM usertbl WHERE userName=?)) ORDER BY tweetId DESC LIMIT 3;", username)

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

func GetFollowersTimedPost(username string, lastTweetId string) (*[]model.Post, *errors.AppError) {

	db := DbConn()
	results, err := db.Query("SELECT tweetId, tweetText, pubTime, userName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.tweetAuthorId=usertbl.userId WHERE tweetAuthorId IN (SELECT followerId FROM followstbl WHERE userId IN (SELECT userId FROM usertbl WHERE userName=?)) AND tweetId<? ORDER BY tweetId DESC LIMIT 3;", username, lastTweetId)

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

func NewPost(newPostData model.NewPost) (bool, *errors.AppError) {
	db := DbConn()

	_, err := db.Exec("INSERT into tweetstbl (tweetText, tweetAuthorId) VALUES (?, (SELECT userId FROM usertbl WHERE userName=?));", newPostData.TweetText, newPostData.UserName)
	if err != nil {
		return false, &errors.AppError{
			Message:    "Error in adding new Tweet",
			StatusCode: http.StatusNotFound,
			Status:     "not found",
		}
	} else {
		return true, nil
	}

}
