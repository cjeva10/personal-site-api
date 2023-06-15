package exchangeHandler

import (
	"github.com/gofiber/fiber/v2"

	"git.chrisevanko.com/personal-site-api.git/database"
	"git.chrisevanko.com/personal-site-api.git/internal/model"
)

func GetExchanges(c *fiber.Ctx) error {
	db := database.DB
	var exchanges []model.Exchange

	db.Find(&exchanges)

	if len(exchanges) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchanges present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Exchanges Found", "data": exchanges})
}

func CreateExchanges(c *fiber.Ctx) error {
	db := database.DB
	exchange := new(model.Exchange)

	err := c.BodyParser(exchange)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

    // check if an exchange with the same symbol has already been created
    var exchangeCheck model.Exchange
    db.Find(&exchangeCheck, "symbol = ?", exchange.Symbol)
    if exchangeCheck.ID != 0 {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Exchange with same symbol already exists", "data": nil})
    }

	err = db.Create(&exchange).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create exchange", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created new exchange", "data": exchange})
}

func GetExchange(c *fiber.Ctx) error {
	db := database.DB
	var exchange model.Exchange

	symbol := c.Params("exchangeSymbol")

	db.Find(&exchange, "symbol = ?", symbol)

	if exchange.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Exchange not found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Exchanges found", "data": exchange})
}

func UpdateExchange(c *fiber.Ctx) error {
	type updateExchange struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Type   string `json:"type"`
	}
	db := database.DB
	var exchange model.Exchange

	// Read the param exchangeId
	id := c.Params("exchangeId")

	// Find the exchange with the given Id
	db.Find(&exchange, "id = ?", id)

	// If no such exchange present return an error
	if exchange.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchange present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateExchangeData updateExchange
	err := c.BodyParser(&updateExchangeData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the exchange
	exchange.Name = updateExchangeData.Name
	exchange.Symbol = updateExchangeData.Symbol
	exchange.Type = updateExchangeData.Type

	// Save the Changes
	db.Save(&exchange)

	// Return the updated exchange
	return c.JSON(fiber.Map{"status": "success", "message": "Exchanges Found", "data": exchange})
}

func DeleteExchange(c *fiber.Ctx) error {
	db := database.DB
	var exchange model.Exchange

	// Read the param exchangeId
	id := c.Params("exchangeId")

	// Find the exchange with the given Id
	db.Find(&exchange, "id = ?", id)

	// If no such exchange present return an error
	if exchange.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No exchange present", "data": nil})
	}

	// Delete the exchange and return error if encountered
	err := db.Delete(&exchange, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete exchange", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Exchange"})
}
