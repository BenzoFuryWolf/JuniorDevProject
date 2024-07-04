package personRoutes

import (
	personHandler "github.com/BenzoFuryWolf/MyProject/internal/handler/person"
	"github.com/gofiber/fiber/v2"
)

func SetupPersonRoutes(router fiber.Router) {
	usePerson := personHandler.Person{}
	person := router.Group("/person")
	//create person
	person.Post("/add", usePerson.CreatePersons)
	//
	person.Get("/pagination/:page/:limit", usePerson.Pagination)
	//read all persons
	person.Get("/all", usePerson.GetPersons)
	//read person
	person.Get("/:personId", usePerson.GetPerson)
	//update person
	person.Put("/:personId", usePerson.UpdatePerson)
	//delete person
	person.Delete("/:personId", usePerson.DeletePerson)
}
