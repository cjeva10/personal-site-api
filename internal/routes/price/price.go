package priceRoutes

import (
	priceHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/price"
	"git.chrisevanko.com/personal-site-api.git/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupPriceRoutes(router fiber.Router) {
	price := router.Group("/price")

	price.Use(middleware.Middleware)

	price.Get("/", priceHandler.GetPrices)
	price.Get("/:assetSymbol", priceHandler.GetPricesByAsset)
	price.Get("/:assetSymbol/:startDate?to:endDate?", priceHandler.GetPricesWithinDateRange)

	price.Put("/:priceId", priceHandler.UpdatePrice)

	price.Delete("/:priceId", priceHandler.DeletePrice)
}
