package responseapi

import (
	"github.com/go-playground/validator/v10"
)

func Error(ve validator.ValidationErrors) []ApiError {
	out := make([]ApiError, len(ve))
	for i, fe := range ve {
		out[i] = ApiError{Code: ValidationError, Title: Title(ValidationError), Message: BindingMessage(fe.Field(), fe.Tag())}
	}
	return out
}
