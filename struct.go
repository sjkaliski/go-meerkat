package meerkat

// Broadcast represents a single broadcast.
type Broadcast struct {
	ID            string      `json:"id"`
	Cover         string      `json:"cover"`
	Broadcaster   *UserInfo   `json:"broadcaster"`
	Caption       string      `json:"caption"`
	Location      string      `json:"location"`
	Place         string      `json:"place"`
	WatchersCount int64       `json:"watchersCount"`
	CommentsCount int64       `json:"commentsCount"`
	LikesCount    int64       `json:"likesCount"`
	RestreamCount int64       `json:"restreamCount"`
	Activities    []*Activity `json:"activities"`
	Influencers   []string    `json:"influencers"`
	Status        string      `json:"status"`
	EndTime       int64       `json:"endTime"`
	TweetID       int64       `json:"tweetId"`
	Tags          []string    `json:"tags"`
}

// User represents a single user with info and stats.
type User struct {
	Info  *UserInfo  `json:"info"`
	Stats *UserStats `json:"stats"`
}

// UserInfo represents user information.
type UserInfo struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	TwitterID   string `json:"twitterId"`
	Privacy     string `json:"privacy"`
	Bio         string `json:"bio"`
}

// UserStats represents user statistics.
type UserStats struct {
	Streams        []Broadcast `json:"streams"`
	StreamsCount   int64       `json:"streamsCount"`
	FollowingCount int64       `json:"followingCount"`
	FollowersCount int64       `json:"followersCount"`
	Score          int64       `json:"score"`
}

// Activity represents an action (comment, restream, like).
type Activity struct {
	Watcher *UserInfo `json:"watcher"`
	Title   string    `json:"title"`
	Message string    `json:"message"`
}
