package domain

import (
	"errors"
	"regexp"
	"strings"
)

// The most inner impl
type User struct {
	Name * string `json:"name";firestore:"name,omitempty"`
	Email * string `json:"email";firestore:"email,omitempty"`
	Id * string `json:"id";firestore:"id,omitempty"`
	ImageSrc * string `json:"imageSrc";firestore:"imageSrc,omitempty"`
}


func isIdValid(id string) bool {
	hasSpaceInMiddleReg := regexp.MustCompile(`[[a-zA-Z0-9@]+[\s]+[a-zA-Z0-9@]+]*`)
	isNotEmptyReg := regexp.MustCompile(`[a-zA-Z0-9@]+`)
	hasSpaceInMiddle := hasSpaceInMiddleReg.MatchString(id)
	isNotEmpty := isNotEmptyReg.MatchString(id)

	return !hasSpaceInMiddle && isNotEmpty
}

func NewUser(Name * string, Email * string, Id * string, ImageSrc * string) (* User, error) {
	id := strings.TrimSpace(*Id)
	valid := isIdValid(id)
	if !valid {
		return nil, errors.New("User id not valid")
	}

	res := &User{Name: Name, Email: Email, Id: &id, ImageSrc: ImageSrc}

	return res, nil

}






