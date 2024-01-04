package entity

import (
	"fmt"
	"net/mail"
	"regexp"
)

var isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString

func validateString(s string, minLen, maxLen int) error {
	if len(s) < minLen || len(s) > maxLen {
		return fmt.Errorf("must contain from %d-%d characters", minLen, maxLen)
	}
	return nil
}

func validateName(s string) error {
	if err := validateString(s, 3, 30); err != nil {
		return err
	}

	if !isValidUsername(s) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func validatePassword(value string) error {
	return validateString(value, 6, 50)
}

func validateEmail(s string) error {
	if err := validateString(s, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(s); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
