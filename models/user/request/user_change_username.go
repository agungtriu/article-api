package request

type UserChangeUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
