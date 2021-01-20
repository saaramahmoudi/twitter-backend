package domain



// The most inner impl
type User struct {
	Name string `json:"name";firestore:"name,omitempty"`
	Email string `json:"email";firestore:"email,omitempty"`
	Id string `json:"id";firestore:"id,omitempty"`
}









