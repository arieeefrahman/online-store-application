package handler

import (
	"online-store-application/database"
	"online-store-application/model/entity"
	"online-store-application/model/request"
	"online-store-application/model/response"
	"online-store-application/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserHandlerRegister(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	if err := user.ValidateUserCreateRequest(); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err,
		})
	}

	var existingUser entity.User
    database.DB.Where("username = ?", user.Username).First(&existingUser)
    if existingUser.ID != uuid.Nil { // Username already exists
        return ctx.Status(400).JSON(fiber.Map{
            "message": "Failed to create user",
			"error": "Username already exists",
        })
    }

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error":   err,
		})
	}

	newUser := entity.User{
		ID:       uuid.New(),
		Name:     user.Name,
		Username: user.Username,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
			"error":   err,
		})
	}

	response := response.UserRegisterResponse{
		ID:        newUser.ID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
		Name:      newUser.Name,
		Username:  newUser.Username,
	}

	return ctx.JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    response,
	})
}
