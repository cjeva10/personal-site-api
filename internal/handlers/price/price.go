package priceHandler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"git.chrisevanko.com/personal-site-api.git/database"
	"git.chrisevanko.com/personal-site-api.git/internal/model"
)

func GetPrices(c *fiber.Ctx) error {
	db := database.DB 
    var prices []model.Price

	db.Find(&prices)

	if len(prices) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No prices present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Prices Found", "data": prices})
}

func GetPricesByAsset(c *fiber.Ctx) error {
	db := database.DB
	var prices []model.Price
    var asset model.Asset

	symbol := c.Params("assetSymbol")

    db.Find(&asset, "symbol = ?", symbol)
    if asset.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found", "data": nil})
    }

	db.Find(&prices, "asset = ?", asset.ID)
	if len(prices) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Prices for this asset not found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Prices found", "data": prices})
}

func UpdatePrice(c *fiber.Ctx) error {
	type updatePrice struct {
		Asset    uint
		Exchange uint 
		Datetime time.Time
		Value    float64
	}

	db := database.DB
	var price model.Price

	id := c.Params("priceId")

	db.Find(&price, "id = ?", id)

	if price.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No price present", "data": nil})
	}

	var updatePriceData updatePrice
	err := c.BodyParser(&updatePriceData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

    price.Asset = updatePriceData.Asset
    price.Exchange = updatePriceData.Exchange
    price.Datetime = updatePriceData.Datetime
    price.Value = updatePriceData.Value

	db.Save(&price)

	return c.JSON(fiber.Map{"status": "success", "message": "Price Updated", "data": price})
}

func DeletePrice(c *fiber.Ctx) error {
	db := database.DB
	var price model.Price

	// Read the param priceId
	id := c.Params("priceId")

	// Find the price with the given Id
	db.Find(&price, "id = ?", id)

	// If no such price present return an error
	if price.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No price present", "data": nil})
	}

	// Delete the price and return error if encountered
	err := db.Delete(&price, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete price", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Price"})
}
