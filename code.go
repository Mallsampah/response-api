package responseapi

const (
	ValidationError = "E001"
)

type Code interface {
	Title(code string) string
}

func Title(code string) string {
	switch code {
	case "E001":
		return "Validation Input Error"
	}

	return ""
}
