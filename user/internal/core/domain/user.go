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

	res := &User{Id: &id, Name: Name, Email: Email, Tag: &tag, ImageSrc: ImageSrc}

	return res, nil

}






