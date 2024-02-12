package handler

import (
	"online-store-application/database"
	"online-store-application/model/entity"
	"online-store-application/model/request"
	"online-store-application/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductCreateHandler(ctx *fiber.Ctx) error {
	product := new(request.Product)

	if err := ctx.BodyParser(product); err != nil {
		return err
	}

	newProduct := entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryID:  product.CategoryID,
	}

	if err := database.DB.Create(&newProduct).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
			"error":   err,
		})
	}

	var category entity.Category
	if err := database.DB.First(&category, newProduct.CategoryID).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to retrieve category data",
			"error":   err,
		})
	}

	response := response.Product{
		ID:           newProduct.ID,
		CreatedAt:    newProduct.CreatedAt,
		UpdatedAt:    newProduct.UpdatedAt,
		DeletedAt:    newProduct.DeletedAt,
		Name:         newProduct.Name,
		Description:  newProduct.Description,
		Price:        newProduct.Price,
		Stock:        newProduct.Stock,
		CategoryName: category.Name,
	}

	return ctx.JSON(fiber.Map{
		"message": "product created successfully",
		"data":    response,
	})
}

func ProductGetAllHandler(ctx *fiber.Ctx) error {
	var products []entity.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to retrieve products",
			"error":   err,
		})
	}

	var responses []response.Product
	for _, product := range products {
		var category entity.Category
		if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to retrieve category data",
				"error":   err,
			})
		}

		response := response.Product{
			ID:           product.ID,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
			DeletedAt:    product.DeletedAt,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: category.Name,
		}
		responses = append(responses, response)
	}

	return ctx.JSON(fiber.Map{
		"message": "products retrieved successfully",
		"data":    responses,
	})
}

func ProductGetByCategoryHandler(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("category_id")
    id, err := strconv.Atoi(categoryId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID or null",
		})
	}
    
	var products []entity.Product
	if err := database.DB.Where("category_id = ?", id).Find(&products).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to retrieve products by category",
			"error":   err,
		})
	}

	var responses []response.Product
	for _, product := range products {
		var category entity.Category
		if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to retrieve category data",
				"error":   err,
			})
		}

		response := response.Product{
			ID:           product.ID,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
			DeletedAt:    product.DeletedAt,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: category.Name,
		}
		responses = append(responses, response)
	}

	return ctx.JSON(fiber.Map{
		"message": "products retrieved successfully by category",
		"data":    responses,
	})
}

func ProductGetByIdHandler(ctx *fiber.Ctx) error {
	productId := ctx.Params("product_id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID or null",
		})
	}

	var product entity.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve product",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product retrieved successfully",
		"data":    product,
	})
}

func ProductDeleteHandler(ctx *fiber.Ctx) error {
	productId := ctx.Params("product_id")
    id, err := strconv.Atoi(productId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID or null",
		})
	}

	if err := database.DB.Where("id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to delete product",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product deleted successfully",
	})
}
