package domain

type Post struct {
	Id * string `json:"id";firestore:"userId,omitempty"`
	UserId * string `json:"userId";firestore:"userId,omitempty"`
	TweetId * string `json:"tweetId";firestore:"tweetId,omitempty"`
	LikedByUserIds []string `json:"likedByUserIds";firestore:"likedByUserIds,omitempty"`
	RetweetedByUserIds []string `json:"retweetedByUserIds";firestore:"retweetedByUserIds,omitempty"`
	MadeAt * int64 `json:"madeAt";firestore:"madeAt,omitempty"`
}

type EventEnum = string

const (
	PostLiked = "PostLiked"
	PostRetweeted = "PostRetweeted"
	PostCreated = "PostCreated"
)


type PostEvent struct {
	Id * string `json:"id"`
	PostId * string `json:"postId"`
	EventType EventEnum `json:"eventType"`
	MadeByUserId * string `json:"madeByUserId"`
	MadeAt * int64 `json:"madeAt"`
	IsReversal * bool `json:"isReversal"`
}


func NewPost(Id * string, UserId * string, TweetId * string, LikedByUserIds []string, RetweetedByUserIds []string, MadeAt * int64) (* Post, error) {

	post := Post{Id: Id, UserId: UserId, TweetId: TweetId, LikedByUserIds: LikedByUserIds, RetweetedByUserIds: RetweetedByUserIds, MadeAt: MadeAt}
	return &post, nil
}

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
func (p * Post) ToggleLikePost(userId * string, MadeAt * int64) (* PostEvent, error){
	index := index(p.LikedByUserIds, *userId)
	isReversal := true
	if index != -1{
		p.LikedByUserIds = append(p.LikedByUserIds[:index], p.LikedByUserIds[index+1:]...)
		return &PostEvent{PostId: p.Id, EventType: PostLiked, MadeByUserId:  userId, MadeAt: MadeAt, IsReversal: &isReversal}, nil
	}else{
		isReversal = false
		p.LikedByUserIds = append(p.LikedByUserIds, *userId)
		return &PostEvent{PostId: p.Id, EventType: PostLiked, MadeByUserId:  userId, MadeAt: MadeAt, IsReversal: &isReversal}, nil
	}
}

func (p * Post) ToggleRetweetPost(userId * string, MadeAt * int64) (* PostEvent, error){
	index := index(p.RetweetedByUserIds, *userId)
	isReversal := true
	if index != -1{
		p.RetweetedByUserIds = append(p.RetweetedByUserIds[:index], p.RetweetedByUserIds[index+1:]...)
		return &PostEvent{PostId: p.Id, EventType: PostRetweeted, MadeByUserId:  userId, MadeAt: MadeAt, IsReversal: &isReversal}, nil
	}else{
		isReversal = false
		p.RetweetedByUserIds = append(p.RetweetedByUserIds, *userId)
		return &PostEvent{PostId: p.Id, EventType: PostRetweeted, MadeByUserId:  userId, MadeAt: MadeAt, IsReversal: &isReversal}, nil
	}
}






