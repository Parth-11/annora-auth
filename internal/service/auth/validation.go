package authservice

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var passwordRegex = regexp.MustCompile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*\W)(?!.* ).{8,16}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	return passwordRegex.MatchString(password)
}
