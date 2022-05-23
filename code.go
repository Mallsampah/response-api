package responseapi

const (
	InvalidInput = "E001"
)

func Title(code string) string {
	switch code {
	case "E001":
		return "Invalid Input"
	}

	return ""
}
