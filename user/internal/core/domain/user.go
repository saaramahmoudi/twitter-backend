package domain

import (
	"errors"
	"regexp"
	"strings"
)

// The most inner impl
type User struct {
	Id		 * string `json:"id";firestore:"id,omitempty"`
	Name     * string `json:"name";firestore:"name,omitempty"`
	Email    * string `json:"email";firestore:"email,omitempty"`
	Tag      * string `json:"tag";firestore:"tag,omitempty"`
	ImageSrc * string `json:"imageSrc";firestore:"imageSrc,omitempty"`
	FollowersId []string `json:"followersId";firestore:"followersId,omitempty"`
	FollowingsId []string `json:"followingsId";firestore:"followingsId,omitempty"`
}


func isIdValid(tag string) bool {
	hasSpaceInMiddleReg := regexp.MustCompile(`[[a-zA-Z0-9@]+[\s]+[a-zA-Z0-9@]+]*`)
	isNotEmptyReg := regexp.MustCompile(`[a-zA-Z0-9@]+`)
	hasSpaceInMiddle := hasSpaceInMiddleReg.MatchString(tag)
	isNotEmpty := isNotEmptyReg.MatchString(tag)

	return !hasSpaceInMiddle && isNotEmpty
}

func NewUser(Name * string, Email * string, Tag * string, ImageSrc * string) (* User, error) {
	tag := strings.TrimSpace(*Tag)
	valid := isIdValid(tag)

	if !valid {
		return nil, errors.New("User tag not valid")
	}

	id := *Email

	res := &User{Id: &id, Name: Name, Email: Email, Tag: &tag, ImageSrc: ImageSrc, FollowersId: []string{}, FollowingsId: []string{}}

	return res, nil
}


func (u *User) followUser(otherUser *User){
	u.FollowingsId = append(u.FollowingsId, *otherUser.Id)
}

func (u *User) unFollowUser(otherUser *User){
	u.FollowingsId = append(u.FollowingsId, *otherUser.Id)
}

func (u *User) getFollower(otherUser *User){
	u.FollowersId = append(u.FollowersId, *otherUser.Id)
}

func (u *User) removeFollower(otherUser *User){
	u.FollowersId = append(u.FollowersId, *otherUser.Id)
}





