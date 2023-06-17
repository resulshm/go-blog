package validation

func ErrorMessage() map[string]string {
	return map[string]string{
		"required": "The field is required",
		"email":    "The field must have a valid email address",
		"min":      "Should be more than the limit",
		"max":      "Should be less than the limit",
	}
}
