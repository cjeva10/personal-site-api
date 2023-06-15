package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    noteRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/note"
    assetRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/asset"
    exchangeRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/exchange"
    priceRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/price"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, world\n")
    })

    api := app.Group("/", logger.New())

    noteRoutes.SetupNoteRoutes(api)

    assetRoutes.SetupAssetRoutes(api)

    exchangeRoutes.SetupExchangeRoutes(api)

    priceRoutes.SetupPriceRoutes(api)
}
