package handler

import (
	"online-store-application/database"
	"online-store-application/model/entity"
	"online-store-application/model/request"
	"online-store-application/model/response"
	"online-store-application/redis"
	"online-store-application/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginReq := new(request.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		return err
	}

	if err := loginReq.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err,
		})
	}

	var existingUser entity.User
	err := database.DB.First(&existingUser, "username = ?", loginReq.Username).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to login",
			"error":   "Username or password is wrong",
		})
	}

	isPasswordValid := util.CheckingPassword(loginReq.Password, existingUser.Password)
	if !isPasswordValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to login",
			"error":   "Username or password is wrong",
		})
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = existingUser.ID
	claims["username"] = existingUser.Username
	claims["name"] = existingUser.Name
	claims["role"] = existingUser.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tkn, err := util.GenerateToken(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to login",
			"error":   err.Error(),
		})
	}

	if err := redis.StoreToken(tkn); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to store token",
			"error":   err.Error(),
		})
	}

	response := response.LoginResponse{
		Token: tkn,
	}

	return ctx.JSON(fiber.Map{
		"message": "Login success",
		"data":    response,
	})
}

func LogoutHandler(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	token := parts[1]

	err := redis.DeleteToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to logout",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "logout success",
	})
}
