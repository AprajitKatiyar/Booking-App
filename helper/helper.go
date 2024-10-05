package helper

import "strings"

func IsValidUserInput(firstName string, lastName string, email string) (bool, bool) {
	isNameValid := len(firstName) > 2 && len(lastName) > 2
	isEmailValid := strings.Contains(email, "@")

	return isNameValid, isEmailValid
}
