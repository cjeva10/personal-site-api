package authRoutes

import (
    "github.com/gofiber/fiber/v2"
    authHandler "git.chrisevanko.com/personal-site-api.git/internal/handlers/auth"
)

func SetupAuthRoutes(router fiber.Router) {
    auth := router.Group("/auth")

    auth.Post("/register", authHandler.Register)
    auth.Post("/login", authHandler.Login)
} 
