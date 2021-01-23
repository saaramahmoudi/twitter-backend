package domain


type Post struct {
	Id * string `json:"id";firestore:"userId,omitempty"`
	UserId * string `json:"userId";firestore:"userId,omitempty"`
	TweetId * string `json:"tweetId";firestore:"tweetId,omitempty"`
	LikedByUserIds []string `json:"likedByUserIds";firestore:"likedByUserIds,omitempty"`
	RetweetedByUserIds []string `json:"retweetedByUserIds";firestore:"retweetedByUserIds,omitempty"`
}





func NewPost(Id * string, UserId * string, TweetId * string, LikedByUserIds []string, RetweetedByUserIds []string) (* Post, error) {
	return &Post{Id: Id, UserId: UserId, TweetId: TweetId, LikedByUserIds: LikedByUserIds, RetweetedByUserIds: RetweetedByUserIds}, nil
}






