package meerkat

// Broadcast represents a single broadcast.
type Broadcast struct {
	ID              string                    `json:"id"`
	Cover           string                    `json:"cover"`
	Broadcaster     *UserInfo                 `json:"broadcaster"`
	Caption         string                    `json:"caption"`
	Summary         string                    `json:"summary"`
	Location        string                    `json:"location"`
	Place           string                    `json:"place"`
	Expiry          int64                     `json:"expiry"`
	WatchersCount   int64                     `json:"watchersCount"`
	CommentsCount   int64                     `json:"commentsCount"`
	LikesCount      int64                     `json:"likesCount"`
	RestreamCount   int64                     `json:"restreamCount"`
	Activities      []*Activity               `json:"activities"`
	Influencers     []string                  `json:"influencers"`
	Status          string                    `json:"status"`
	EndTime         int64                     `json:"endTime"`
	TweetID         int64                     `json:"tweetId"`
	Tags            []string                  `json:"tags"`
	FollowUpActions *BroadcastFollowUpActions `json:"followupActions"`
}

// User represents a single user with info and stats.
type User struct {
	Info                *UserInfo            `json:"info"`
	Stats               *UserStats           `json:"stats"`
	UserFollowUpActions *UserFollowUpActions `json:"followupActions"`
}

// UserInfo represents user information.
type UserInfo struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	DisplayName   string `json:"displayName"`
	TwitterID     string `json:"twitterId"`
	Privacy       string `json:"privacy"`
	Bio           string `json:"bio"`
	Follow        string `json:"follow"`
	Profile       string `json:"profile"`
	PhotoUrl      string `json:"photoUrl"`
	ThumbImageUrl string `json:"profileThumbImage"`
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

// BroadcastFollowUpActions represents follow up actions
// for a broadcast request.
type BroadcastFollowUpActions struct {
	Likes     string `json:"likes"`
	Delete    string `json:"delete"`
	Playlist  string `json:"playlist"`
	Restreams string `json:"restreams"`
	Comments  string `json:"comments"`
	Watchers  string `json:"watchers"`
}

// UserFollowUpActions represents follow up actions
// for a user request.
type UserFollowUpActions struct {
	ProfileThumbImage string `json:"profileThumbImage"`
	ProfileImage      string `json:"profileImage"`
	Followers         string `json:"followers"`
	Following         string `json:"following"`
	Report            string `json:"report"`
	Follow            string `json:"follow"`
}
