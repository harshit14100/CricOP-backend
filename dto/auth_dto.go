package dto

type SignupRequest struct {
	Name     string `json:"name"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
