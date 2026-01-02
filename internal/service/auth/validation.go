package authservice

import (
	"regexp"
	"unicode"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 || regexp.MustCompile(`\s`).MatchString(password) {
		return false
	}

	var hasNumber, hasUpperCase, hasLowerCase, hasSpecial bool
	for _, c := range password {
		switch {
			case unicode.IsNumber(c):
				hasNumber = true
			case unicode.IsUpper(c):
				hasUpperCase = true
			case unicode.IsLower(c):
				hasLowerCase = true
			case unicode.IsPunct(c) || unicode.IsSymbol(c):
				hasSpecial = true
		}
	}

	return hasNumber && hasUpperCase && hasLowerCase && hasSpecial
}
