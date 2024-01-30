package router

import (
	personRoutes "github.com/BenzoFuryWolf/MyProject/internal/routes"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api", logger.New())
	personRoutes.SetupPersonRoutes(api)
}
