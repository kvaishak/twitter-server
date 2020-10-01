package dao

import (
	"net/http"

	"github.com/kvaishak/twitter-server/errors"
	"github.com/kvaishak/twitter-server/model"
)

func GetFollowersPost(username string) (*[]model.Post, *errors.AppError) {

	db := DbConn()

	results, err := db.Query("SELECT TweetId, TweetText, PubTime, UserName, FirstName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.TweetAuthorID=usertbl.UserId WHERE TweetAuthorID IN (SELECT FollowerId FROM followstbl WHERE UserId IN (SELECT UserId FROM usertbl WHERE UserName=?)) ORDER BY TweetId DESC LIMIT 3;", username)

	defer db.Close()

	if results != nil {
		//Getting all Followers Post data
		postsArr := []model.Post{}
		for results.Next() {
			var post model.Post
			err = results.Scan(&post.TweetID, &post.TweetText, &post.PublishTime, &post.UserName, &post.FirstName)
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
	results, err := db.Query("SELECT TweetId, TweetText, PubTime, UserName, FirstName FROM tweetstbl INNER JOIN usertbl ON tweetstbl.TweetAuthorID=usertbl.UserId WHERE TweetAuthorID IN (SELECT FollowerId FROM followstbl WHERE UserId IN (SELECT UserId FROM usertbl WHERE UserName=?)) AND TweetId<? ORDER BY TweetId DESC LIMIT 3;", username, lastTweetId)

	defer db.Close()

	if results != nil {
		//Getting all Followers Post data
		postsArr := []model.Post{}
		for results.Next() {
			var post model.Post
			err = results.Scan(&post.TweetID, &post.TweetText, &post.PublishTime, &post.UserName, &post.FirstName)
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

	_, err := db.Exec("INSERT into tweetstbl (TweetText, TweetAuthorID) VALUES (?, (SELECT UserId FROM usertbl WHERE UserName=?));", newPostData.TweetText, newPostData.UserName)
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
