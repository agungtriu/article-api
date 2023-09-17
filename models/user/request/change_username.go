package request

type ChangeUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
