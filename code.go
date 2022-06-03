package responseapi

const (
	InvalidInput = "E001"
	AlreadyExist = "E002"
	NotExist     = "E003"

	Unauthorized = "A001"
)

func Title(code string) string {
	switch code {
	case "E001":
		return "Invalid Input"
	case "E002":
		return "Already Exists"
	case "E003":
		return "Data Not Exists"
	case "A001":
		return "Unauthorized"
	}

	return ""
}
