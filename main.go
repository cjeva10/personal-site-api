package main

import (
	"github.com/gofiber/fiber/v2"
    "git.chrisevanko.com/personal-site-api.git/database"
    "git.chrisevanko.com/personal-site-api.git/router"
)

func main() {
	app := fiber.New()

    database.ConnectDB()

    router.SetupRoutes(app)

    app.Listen(":3000")
}
