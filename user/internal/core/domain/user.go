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

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func ToggleList(instance string, list * []string){
	index := index(*list, instance)
	if index != -1{
		*list = append((*list)[:index], (*list)[index+1:]...)
	}else{
		*list = append(*list, instance)
	}
}
func (u *User) TogglefollowingUser(otherUser *User){
	ToggleList(*otherUser.Id, &u.FollowingsId)
}


func (u *User) ToggleFollower(otherUser *User){
	ToggleList(*otherUser.Id, &u.FollowersId)
}





