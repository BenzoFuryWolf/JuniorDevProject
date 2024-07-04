package main

import (
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/config"
	"github.com/BenzoFuryWolf/MyProject/database"
	"github.com/BenzoFuryWolf/MyProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	database.ConnectDB()

	PORT := fmt.Sprintf(":%s", config.Config("PORT"))
	router.SetupRoutes(app)

	err := app.Listen(PORT)
	if err != nil {
		return
	}
}
