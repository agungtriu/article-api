package request

type UserChangeEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
