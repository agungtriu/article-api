package base

type ErrorResponse struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
}
