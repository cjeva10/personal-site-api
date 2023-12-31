package router

import (
	assetRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/asset"
	exchangeRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/exchange"
	noteRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/note"
	priceRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/price"
    authRoutes "git.chrisevanko.com/personal-site-api.git/internal/routes/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, world\n")
    })

    api := app.Group("/", logger.New())

    authRoutes.SetupAuthRoutes(api)

    noteRoutes.SetupNoteRoutes(api)

    assetRoutes.SetupAssetRoutes(api)

    exchangeRoutes.SetupExchangeRoutes(api)

    priceRoutes.SetupPriceRoutes(api)
}
