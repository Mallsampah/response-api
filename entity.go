package responseapi

type ApiError struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type ApiSuccess[T any, R any] struct {
	Data Data[T, R] `json:"data"`
}

type Data[T any, R any] struct {
	ID           string `json:"id"`
	Attributes   T      `json:"attributes"`
	Relationship R      `json:"relationship"`
}
