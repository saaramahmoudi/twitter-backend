package domain

type Post struct {
	Id * string `json:"id";firestore:"userId,omitempty"`
	UserId * string `json:"userId";firestore:"userId,omitempty"`
	TweetId * string `json:"tweetId";firestore:"tweetId,omitempty"`
	LikedByUserIds []string `json:"likedByUserIds";firestore:"likedByUserIds,omitempty"`
	RetweetedByUserIds []string `json:"retweetedByUserIds";firestore:"retweetedByUserIds,omitempty"`
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
}


func NewPost(Id * string, UserId * string, TweetId * string, LikedByUserIds []string, RetweetedByUserIds []string) (* PostEvent, error) {
	post := Post{Id: Id, UserId: UserId, TweetId: TweetId, LikedByUserIds: LikedByUserIds, RetweetedByUserIds: RetweetedByUserIds}
	return &PostEvent{Post: &post, EventType: PostPublished, MadeByUserId: UserId}, nil
}

func (p * Post) LikePost(userId * string) (* PostEvent, error){
	p.LikedByUserIds = append(p.LikedByUserIds, *userId)
	return &PostEvent{Post: p, EventType: PostLiked, MadeByUserId:  userId}, nil
}

func (p * Post) RetweetPost(userId * string) (* PostEvent, error){
	p.LikedByUserIds = append(p.LikedByUserIds, *userId)
	return &PostEvent{Post: p, EventType: PostRetweeted, MadeByUserId:  userId}, nil
}






