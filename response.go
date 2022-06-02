package responseapi

import (
	"github.com/go-playground/validator/v10"
)

func ValidationError(ve validator.ValidationErrors) []ApiError {
	out := make([]ApiError, len(ve))
	for i, fe := range ve {
		out[i] = ApiError{Code: InvalidInput, Title: Title(InvalidInput), Message: BindingMessage(fe.Field(), fe.Tag())}
	}
	return out
}

func Error(code string, message string) ApiError {
	return ApiError{Code: code, Title: Title(code), Message: message}
}

func Success(id string, T any) (apiSuccess ApiSuccess[any]) {
	return ApiSuccess[any]{Data: Data[any]{ID: id, Attributes: Attributes[any]{Attributes: T}}}
}
