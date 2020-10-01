package model

type Post struct {
	TweetID     int    `json:"tweetid"`
	TweetText   string `json:"tweet_text"`
	PublishTime string `json:"time"`
	UserName    string `json:"username"`
	FirstName   string `json:"firstname"`
}

type NewPost struct {
	TweetText string `json:"tweetText"`
	UserName  string `json:"username"`
}
