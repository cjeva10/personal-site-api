package assetRoutes

import (
    "github.com/gofiber/fiber/v2"
    exchangeHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/exchange"
)

func SetupExchangeRoutes(router fiber.Router) {
    exchange := router.Group("/exchange")

    exchange.Post("/", exchangeHandler.CreateExchanges)

    exchange.Get("/", exchangeHandler.GetExchanges)
    
    exchange.Get("/:exchangeSymbol", exchangeHandler.GetExchange)

    exchange.Put("/:exchangeId", exchangeHandler.UpdateExchange)

    exchange.Delete("/:exchangeId", exchangeHandler.DeleteExchange)
}
