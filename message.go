package responseapi

func BindingMessage(field string, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return "Invalid email"
	case "e164":
		return "Invalid format phone number"
	}
	return ""
}
