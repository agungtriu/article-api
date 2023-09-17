package request

type ChangeProfile struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
