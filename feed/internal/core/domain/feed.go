package domain







type Feed struct {
	Id * string
	UserId * string
	PostIds [] string
}

type FeedProvider interface {
	GetFeed() * Feed
}
type PersonalFeed struct {
	Feed
}
func (pf * PersonalFeed) GetFeed() * Feed {

	return &pf.Feed
}

type PublicPageFeed struct {
	Feed
}

func (pf * PublicPageFeed) GetFeed() * Feed {

	return &pf.Feed
}




























