package request

type ChangeEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
