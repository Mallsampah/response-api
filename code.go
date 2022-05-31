package responseapi

const (
	InvalidInput = "E001"
	AlreadyExist = "E002"
)

func Title(code string) string {
	switch code {
	case "E001":
		return "Invalid Input"
	case "E002":
		return "Already Exist"
	}

	return ""
}
