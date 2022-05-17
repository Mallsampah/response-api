package responseapi

type ApiError struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
