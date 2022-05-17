package responseapi

const (
	ValidationError = "E001"
)

func Title(code string) string {
	switch code {
	case "E001":
		return "Validation Input Error"
	}

	return ""
}
