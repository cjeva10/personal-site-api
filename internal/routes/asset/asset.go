package assetRoutes

import (
	assetHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/asset"
	"git.chrisevanko.com/personal-site-api.git/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupAssetRoutes(router fiber.Router) {
	asset := router.Group("/asset")

	asset.Get("/", assetHandler.GetAssets)
	asset.Get("/:assetSymbol", assetHandler.GetAsset)

    // authenticate all routes that can change the db
	asset.Use(middleware.Middleware)

	asset.Post("/", assetHandler.CreateAssets)

	asset.Put("/:assetId", assetHandler.UpdateAsset)

	asset.Delete("/:assetId", assetHandler.DeleteAsset)
}
