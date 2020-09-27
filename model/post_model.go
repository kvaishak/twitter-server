package model

type Post struct {
	TweetText   string `json:"tweet_text"`
	PublishTime string `json:"time"`
	UserName    string `json:"username"`
}
