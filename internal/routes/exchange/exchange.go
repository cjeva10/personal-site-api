package assetRoutes

import (
	exchangeHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/exchange"
	"git.chrisevanko.com/personal-site-api.git/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupExchangeRoutes(router fiber.Router) {
	exchange := router.Group("/exchange")

	exchange.Get("/", exchangeHandler.GetExchanges)

	exchange.Get("/:exchangeSymbol", exchangeHandler.GetExchange)

	exchange.Use(middleware.Middleware)

	exchange.Post("/", exchangeHandler.CreateExchanges)

	exchange.Put("/:exchangeId", exchangeHandler.UpdateExchange)

	exchange.Delete("/:exchangeId", exchangeHandler.DeleteExchange)
}
