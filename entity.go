package responseapi

type ApiError struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type ApiSuccess[T any] struct {
	Data Data[T] `json:"data"`
}

type Data[T any] struct {
	ID         string        `json:"id"`
	Attributes Attributes[T] `json:"attributes"`
}

type Attributes[T any] struct {
	Attributes T
}
