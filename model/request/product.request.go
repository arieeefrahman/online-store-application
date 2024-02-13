package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

func (req *ProductRequest) Validate() map[string]string {
	validate := validator.New()
	err := validate.Struct(req)

	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := fieldErr.Field()
		tag := fieldErr.Tag()

		switch tag {
		case "required":
			validationErrors[field] = "This field is required."
		default:
			validationErrors[field] = fmt.Sprintf("'%s' is invalid for field '%s'.", fieldErr.Value(), field)
		}
	}

	return validationErrors
}
