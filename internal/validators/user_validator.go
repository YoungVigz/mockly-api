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

func CreateUserValidator(userRequest *models.UserCreateRequest) ([]string, error) {
	var validationErrors []string

	if userRequest.Username == "" {
		validationErrors = append(validationErrors, "Username is required")
	} else {
		if len(userRequest.Username) < 3 {
			validationErrors = append(validationErrors, "Username must be at least 3 characters long")
		}
		if len(userRequest.Username) > 20 {
			validationErrors = append(validationErrors, "Username cannot be longer than 20 characters")
		}
		if !usernameRegex.MatchString(userRequest.Username) {
			validationErrors = append(validationErrors, "Username can only contain letters, numbers, underscores and hyphens")
		}
	}

	if userRequest.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	} else if !emailRegex.MatchString(userRequest.Email) {
		validationErrors = append(validationErrors, "Invalid email format. Example: user@example.com")
	}

	if userRequest.Password == "" {
		validationErrors = append(validationErrors, "Password is required")
	} else {
		var passwordErrors []string

		if len(userRequest.Password) < 8 {
			passwordErrors = append(passwordErrors, "at least 8 characters")
		}
		if !regexp.MustCompile(`[A-Z]`).MatchString(userRequest.Password) {
			passwordErrors = append(passwordErrors, "one uppercase letter")
		}
		if !regexp.MustCompile(`[a-z]`).MatchString(userRequest.Password) {
			passwordErrors = append(passwordErrors, "one lowercase letter")
		}
		if !regexp.MustCompile(`\d`).MatchString(userRequest.Password) {
			passwordErrors = append(passwordErrors, "one number")
		}
		if !regexp.MustCompile(`[@$!%*?&]`).MatchString(userRequest.Password) {
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
