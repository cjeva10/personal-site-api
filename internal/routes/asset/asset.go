package assetRoutes

import (
    "github.com/gofiber/fiber/v2"
    assetHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/asset"
)

func SetupAssetRoutes(router fiber.Router) {
    asset := router.Group("/asset")

    asset.Post("/", assetHandler.CreateAssets)

    asset.Get("/", assetHandler.GetAssets)
    
    asset.Get("/:assetSymbol", assetHandler.GetAsset)

    asset.Put("/:assetId", assetHandler.UpdateAsset)

    asset.Delete("/:assetId", assetHandler.DeleteAsset)
}
