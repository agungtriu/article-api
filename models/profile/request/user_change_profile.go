package request

type UserChangeProfile struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
