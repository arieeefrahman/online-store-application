package request

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

func (req *UserCreateRequest) Validate() map[string]string {
	validate := validator.New()
	err := validate.Struct(req)

	if err == nil {
		return nil
	}

	// Create a map for error translations
	validationErrors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := fieldErr.Field()
		tag := fieldErr.Tag()

		switch tag {
		case "required":
			validationErrors[field] = "This field is required."
		case "min":
			validationErrors[field] = fmt.Sprintf("This field must be at least %v characters long.", fieldErr.Param())
		default:
			validationErrors[field] = fmt.Sprintf("'%s' is invalid for field '%s'.", fieldErr.Value(), field)
		}
	}

	var validUsernamePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validUsernamePattern.MatchString(req.Username) {
        validationErrors["Username"] = "Username can only contain letters, numbers, underscores, and hyphens."
    }

	return validationErrors
}
