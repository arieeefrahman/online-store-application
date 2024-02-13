package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CartItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
}

func (req *CartItemRequest) Validate() map[string]string {
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