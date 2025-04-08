package validators

import (
	"errors"
	"regexp"
	"strings"

	"github.com/YoungVigz/mockly-api/internal/models"
)

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
)

func CreateUserValidator(password *models.UserAuthRequest) ([]string, error) {
	var validationErrors []string

	if password.Username == "" {
		validationErrors = append(validationErrors, "Username is required")
	} else {
		if len(password.Username) < 3 {
			validationErrors = append(validationErrors, "Username must be at least 3 characters long")
		}
		if len(password.Username) > 20 {
			validationErrors = append(validationErrors, "Username cannot be longer than 20 characters")
		}
		if !usernameRegex.MatchString(password.Username) {
			validationErrors = append(validationErrors, "Username can only contain letters, numbers, underscores and hyphens")
		}
	}

	if password.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	} else if !emailRegex.MatchString(password.Email) {
		validationErrors = append(validationErrors, "Invalid email format. Example: user@example.com")
	}

	if password.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	} else {
		var passwordErrors []string

		if len(password.Password) < 8 {
			passwordErrors = append(passwordErrors, "at least 8 characters")
		}
		if !regexp.MustCompile(`[A-Z]`).MatchString(password.Password) {
			passwordErrors = append(passwordErrors, "one uppercase letter")
		}
		if !regexp.MustCompile(`[a-z]`).MatchString(password.Password) {
			passwordErrors = append(passwordErrors, "one lowercase letter")
		}
		if !regexp.MustCompile(`\d`).MatchString(password.Password) {
			passwordErrors = append(passwordErrors, "one number")
		}
		if !regexp.MustCompile(`[@$!%*?&]`).MatchString(password.Password) {
			passwordErrors = append(passwordErrors, "one special character (@$!%*?&)")
		}

		if len(passwordErrors) > 0 {
			validationErrors = append(validationErrors,
				"Password must contain: "+strings.Join(passwordErrors, ", "))
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors, errors.New("validation failed")
	}

	return nil, nil
}

func LoginValidator(password *models.UserLoginRequest) ([]string, error) {

	var validationErrors []string

	if password.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	} else if !emailRegex.MatchString(password.Email) {
		validationErrors = append(validationErrors, "Invalid email format. Example: user@example.com")
	}

	if password.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	}

	if len(validationErrors) > 0 {
		return validationErrors, errors.New("validation failed")
	}

	return nil, nil
}

func PasswordValidator(password string) ([]string, error) {
	var validationErrors []string

	if password == "" {
		validationErrors = append(validationErrors, "Password is required")
	} else {
		var passwordErrors []string

		if len(password) < 8 {
			passwordErrors = append(passwordErrors, "at least 8 characters")
		}
		if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
			passwordErrors = append(passwordErrors, "one uppercase letter")
		}
		if !regexp.MustCompile(`[a-z]`).MatchString(password) {
			passwordErrors = append(passwordErrors, "one lowercase letter")
		}
		if !regexp.MustCompile(`\d`).MatchString(password) {
			passwordErrors = append(passwordErrors, "one number")
		}
		if !regexp.MustCompile(`[@$!%*?&]`).MatchString(password) {
			passwordErrors = append(passwordErrors, "one special character (@$!%*?&)")
		}

		if len(passwordErrors) > 0 {
			validationErrors = append(validationErrors,
				"Password must contain: "+strings.Join(passwordErrors, ", "))
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors, errors.New("validation failed")
	}

	return nil, nil
}
