package auth

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

var usernameRe = regexp.MustCompile(`^[a-z0-9_]{3,20}$`)

func NameValidate(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return name
	}

	return name
}

func EmailValidate(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
func UsernameValidate(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func ValidateUsername(username string) error {
	if !usernameRe.MatchString(username) {
		return fmt.Errorf("username must be 3-20 characters and contain only lowercase letters, numbers, or underscores")
	}
	return nil
}
func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

func ValidateName(name string) error {
	if len(strings.TrimSpace(name)) > 100 {
		return fmt.Errorf("name is too long")
	}
	return nil
}
