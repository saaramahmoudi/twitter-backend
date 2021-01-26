package domain

import "time"

type Post struct {
	Id * string `json:"id";firestore:"userId,omitempty"`
	UserId * string `json:"userId";firestore:"userId,omitempty"`
	TweetId * string `json:"tweetId";firestore:"tweetId,omitempty"`
	LikedByUserIds []string `json:"likedByUserIds";firestore:"likedByUserIds,omitempty"`
	RetweetedByUserIds []string `json:"retweetedByUserIds";firestore:"retweetedByUserIds,omitempty"`
	MadeAt * time.Time
}

type EventEnum = string

const (
	PostPublished = "PostPublished"
	PostLiked = "PostLiked"
	PostRetweeted = "PostRetweeted"
)


type PostEvent struct {
	Post * Post
	EventType EventEnum
	MadeByUserId * string
	MadeAt * time.Time
}


func NewPost(Id * string, UserId * string, TweetId * string, LikedByUserIds []string, RetweetedByUserIds []string, MadeAt * time.Time) (* PostEvent, error) {
	post := Post{Id: Id, UserId: UserId, TweetId: TweetId, LikedByUserIds: LikedByUserIds, RetweetedByUserIds: RetweetedByUserIds, MadeAt: MadeAt}
	return &PostEvent{Post: &post, EventType: PostPublished, MadeByUserId: UserId, MadeAt: MadeAt}, nil
}

func (p * Post) ToggleLikePost(userId * string, MadeAt * time.Time) (* PostEvent, error){
	p.LikedByUserIds = append(p.LikedByUserIds, *userId)
	return &PostEvent{Post: p, EventType: PostLiked, MadeByUserId:  userId, MadeAt: MadeAt}, nil
}

func (p * Post) ToggleRetweetPost(userId * string, MadeAt * time.Time) (* PostEvent, error){
	p.LikedByUserIds = append(p.LikedByUserIds, *userId)
	return &PostEvent{Post: p, EventType: PostRetweeted, MadeByUserId:  userId, MadeAt: MadeAt}, nil
}






