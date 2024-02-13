package handler

import (
	"log"
	"online-store-application/database"
	"online-store-application/model/entity"
	"online-store-application/model/request"
	"online-store-application/model/response"
	"online-store-application/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CartItemCreateHandler(ctx *fiber.Ctx) error {
	cartItemReq := new(request.CartItemRequest)
	if err := ctx.BodyParser(cartItemReq); err != nil {
		return err
	}

	if err := cartItemReq.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err,
		})
	}

	authHeader := ctx.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	token := parts[1]
	claims, err := util.DecodeToken(token)
	if err != nil {
		log.Fatalf("Error decoding token: %v", err)
	}

	user_id, _ := claims["user_id"].(string)
	username, _ := claims["username"].(string)

	var product entity.Product
	if err := database.DB.First(&product, cartItemReq.ProductID).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	newCartItem := entity.CartItem{
		ProductID: cartItemReq.ProductID,
		Quantity:  cartItemReq.Quantity,
		UserID:    user_id,
	}

	newCartItem.Product = product
	price := newCartItem.Quantity * product.Price
	newCartItem.Price = price

	var existingCartItem entity.CartItem
	if err := database.DB.Where("user_id = ? AND product_id = ?", newCartItem.UserID, cartItemReq.ProductID).First(&existingCartItem).Error; err == nil {
		existingCartItem.Quantity += cartItemReq.Quantity
		existingCartItem.Price = existingCartItem.Quantity * product.Price

		if err := database.DB.Save(&existingCartItem).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to update cart item",
				"error":   err,
			})
		}

		response := response.CartItemResponse{
			ID:          existingCartItem.ID,
			CreatedAt:   existingCartItem.CreatedAt,
			UpdatedAt:   existingCartItem.UpdatedAt,
			ProductID:   existingCartItem.Product.ID,
			ProductName: existingCartItem.Product.Name,
			Quantity:    existingCartItem.Quantity,
			Price:       existingCartItem.Price,
			Username:    username,
		}

		return ctx.JSON(fiber.Map{
			"message": "cart item updated successfully",
			"data":    response,
		})
	}

	if err := database.DB.Create(&newCartItem).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
			"error":   err,
		})
	}

	response := response.CartItemResponse{
		ID:          newCartItem.ID,
		CreatedAt:   newCartItem.CreatedAt,
		UpdatedAt:   newCartItem.UpdatedAt,
		ProductID:   newCartItem.Product.ID,
		ProductName: newCartItem.Product.Name,
		Quantity:    newCartItem.Quantity,
		Price:       newCartItem.Price,
		Username:    username,
	}

	return ctx.JSON(fiber.Map{
		"message": "product added to cart successfully",
		"data":    response,
	})
}

func CartItemUpdateHandler(ctx *fiber.Ctx) error {
	cartItemID := ctx.Params("cart_item_id")

	cartItemUpdate := new(request.CartItemRequest)
	if err := ctx.BodyParser(cartItemUpdate); err != nil {
		return err
	}

	var cartItem entity.CartItem
	if err := database.DB.First(&cartItem, cartItemID).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "cart item not found",
		})
	}

	var product entity.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	cartItem.Quantity = cartItemUpdate.Quantity
	price := cartItem.Quantity * product.Price
	cartItem.Price = cartItemUpdate.Quantity * price

	if err := database.DB.Save(&cartItem).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to update cart item",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "cart item updated successfully",
		"data":    cartItem,
	})
}

func CartItemDeleteHandler(ctx *fiber.Ctx) error {
	cartItemID := ctx.Params("cart_item_id")

	if err := database.DB.Delete(&entity.CartItem{}, cartItemID).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to delete cart item",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "cart item deleted successfully",
	})
}

func CartItemGetByUserIdHandler(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	token := parts[1]
	claims, err := util.DecodeToken(token)

	if err != nil {
		log.Fatalf("Error decoding token: %v", err)
	}

	user_id, _ := claims["user_id"].(string)
	username, _ := claims["username"].(string)

	var cartItems []entity.CartItem
	if err := database.DB.Where("user_id = ?", user_id).Find(&cartItems).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to retrieve cart items",
			"error":   err.Error(),
		})
	}

	var responseItems []response.CartItemResponse
	var totalPrice int
	for _, cartItem := range cartItems {
		var product entity.Product
		if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"message": "product not found",
			})
		}

		cartItem.Product = product
		responseItems = append(responseItems, response.CartItemResponse{
			ID:          cartItem.ID,
			CreatedAt:   cartItem.CreatedAt,
			UpdatedAt:   cartItem.UpdatedAt,
			ProductID:   cartItem.Product.ID,
			ProductName: cartItem.Product.Name,
			Quantity:    cartItem.Quantity,
			Price:       cartItem.Price,
			Username:    username,
		})
		totalPrice += cartItem.Price
	}

	return ctx.JSON(fiber.Map{
		"message":     "cart items retrieved successfully",
		"data":        responseItems,
		"total_price": totalPrice,
	})
}
