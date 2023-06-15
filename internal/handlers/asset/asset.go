package assetHandler

import (
	"github.com/gofiber/fiber/v2"

	"git.chrisevanko.com/personal-site-api.git/database"
	"git.chrisevanko.com/personal-site-api.git/internal/model"
)

func GetAssets(c *fiber.Ctx) error {
	db := database.DB
	var assets []model.Asset

	db.Find(&assets)

	if len(assets) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No assets present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Assets Found", "data": assets})
}

func CreateAssets(c *fiber.Ctx) error {
	db := database.DB
	asset := new(model.Asset)

	err := c.BodyParser(asset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// check if an asset with the same symbol has already been created
	var assetCheck model.Asset
	db.Find(&assetCheck, "symbol = ?", asset.Symbol)
	if assetCheck.ID != 0 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Asset with same symbol already exists", "data": nil})
	}

	err = db.Create(&asset).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create asset", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created new asset", "data": asset})
}

func GetAsset(c *fiber.Ctx) error {
	db := database.DB
	var asset model.Asset

	symbol := c.Params("assetSymbol")

	db.Find(&asset, "symbol = ?", symbol)

	if asset.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Asset not found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Assets found", "data": asset})
}

func UpdateAsset(c *fiber.Ctx) error {
	type updateAsset struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Type   string `json:"type"`
	}
	db := database.DB
	var asset model.Asset

	// Read the param assetId
	id := c.Params("assetId")

	// Find the asset with the given Id
	db.Find(&asset, "id = ?", id)

	// If no such asset present return an error
	if asset.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No asset present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateAssetData updateAsset
	err := c.BodyParser(&updateAssetData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the asset
	asset.Name = updateAssetData.Name
	asset.Symbol = updateAssetData.Symbol
	asset.Type = updateAssetData.Type

	// Save the Changes
	db.Save(&asset)

	// Return the updated asset
	return c.JSON(fiber.Map{"status": "success", "message": "Assets Found", "data": asset})
}

func DeleteAsset(c *fiber.Ctx) error {
	db := database.DB
	var asset model.Asset

	// Read the param assetId
	id := c.Params("assetId")

	// Find the asset with the given Id
	db.Find(&asset, "id = ?", id)

	// If no such asset present return an error
	if asset.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No asset present", "data": nil})
	}

	// Delete the asset and return error if encountered
	err := db.Delete(&asset, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete asset", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Asset"})
}
