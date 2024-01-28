package main

import (
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/config"
	"github.com/BenzoFuryWolf/MyProject/database"
	"github.com/BenzoFuryWolf/MyProject/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	PORT := fmt.Sprintf(":%s", config.Config("PORT"))
	router.SetupRoutes(app)
	app.Listen(PORT)
}
