package is_valid

import "regexp"

func IsValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return false
	}
	return matched
}
