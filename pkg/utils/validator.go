package utils

import "regexp"

func ValidateStringNumber(s string) bool {
	phoneRegex := regexp.MustCompile("^\\+?[0-9]{11,13}$")
	if !phoneRegex.MatchString(s) {
		return false
	}
	return true
}
func ValidateMail(s string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	if !emailRegex.MatchString(s) {
		return false
	}
	return true
}
