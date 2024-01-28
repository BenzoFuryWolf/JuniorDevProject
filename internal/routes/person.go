package personRoutes

import (
	personHandler "github.com/BenzoFuryWolf/MyProject/internal/handler/person"
	"github.com/gofiber/fiber/v2"
)

func SetupPersonRoutes(router fiber.Router) {
	person := router.Group("/person")
	//create person
	person.Post("/add", personHandler.CreatePersons)
	//read all persons
	person.Get("/all", personHandler.GetPersons)
	//read person
	person.Get("/:personId", personHandler.GetPerson)
	//update person
	person.Put("/:personId", personHandler.UpdatePerson)
	//delete person
	person.Delete("/:personId", personHandler.DeletePerson)
}
