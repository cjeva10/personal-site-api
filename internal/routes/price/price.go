package priceRoutes

import (
	priceHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/price"
	"git.chrisevanko.com/personal-site-api.git/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func SetupPriceRoutes(router fiber.Router) {
	price := router.Group("/price")

	// all price endpoints are protected
	price.Use(keyauth.New(keyauth.Config{
		KeyLookup: "header:key",
		Validator: middleware.ValidateApiKey,
	}))

	price.Get("/", priceHandler.GetPrices)

	price.Get("/:assetSymbol", priceHandler.GetPricesByAsset)
	price.Get("/:assetSymbol/:startDate?to:endDate?", priceHandler.GetPricesWithinDateRange)

	price.Put("/:priceId", priceHandler.UpdatePrice)

	price.Delete("/:priceId", priceHandler.DeletePrice)
}
