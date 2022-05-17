package responseapi

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

type ResponseError interface {
	GetError(err error) []ApiError
}

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return " "
}

func GetError(err error) []ApiError {
	fmt.Println("okeokeoke")
	var ve validator.ValidationErrors
	out := make([]ApiError, len(ve))
	if errors.As(err, &ve) {
		for i, fe := range ve {
			out[i] = ApiError{Field: fe.Field(), Message: MsgForTag(fe.Tag())}
		}
	}
	return out
}
