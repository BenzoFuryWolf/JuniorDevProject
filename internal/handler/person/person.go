package personHandler

import (
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/database"
	"github.com/BenzoFuryWolf/MyProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
)

type updatePerson struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationalize string `json:"nationalize"`
}

type Person struct {
	persons  []model.Person_info
	person   model.Person_info
	upPerson updatePerson
}

type Personer interface {
	GetPersons(c *fiber.Ctx) error
	CreatePersons(c *fiber.Ctx) error
	UpdatePerson(c *fiber.Ctx) error
	DeletePerson(c *fiber.Ctx) error
	GetPerson(c *fiber.Ctx) error
	Pagination(c *fiber.Ctx) error
}

func (p Person) GetPersons(c *fiber.Ctx) error {
	var db = database.DB
	var person = p.persons
	db.Find(&person)

	if len(person) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No persons present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Persons found", "data": person})
}

func (p Person) CreatePersons(c *fiber.Ctx) error {
	c.GetRespHeader("Access-Control-Allow-Origin", "*")
	var api PersonEnricher
	api = API{}
	var db = database.DB
	var person = p.person
	err := c.BodyParser(&person)
	if err != nil {
		return c.JSON(fiber.Map{"status": "field", "message": "Create user has failed", "data": person})
	}

	nazionalize := api.nazy(person.Name)

	age, err := strconv.Atoi(api.myAge(person.Name))
	if err != nil {
		return err
	}

	gender := api.myGender(person.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	person.ID = uuid.New()
	person.Age = age
	person.Gender = gender
	person.Nationalize = nazionalize
	err = db.Create(&person).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create person", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Create Person", "data": person})
}

func (p Person) UpdatePerson(c *fiber.Ctx) error {
	c.GetRespHeader("Access-Control-Allow-Origin", "*")
	db := database.DB
	var person = p.person

	id := c.Params("personId")
	db.Find(&person, "id = ?", id)

	if person.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No person present", "data": nil})
	}

	var updatePersonData = p.upPerson
	err := c.BodyParser(&updatePersonData)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	person.Name = updatePersonData.Name
	person.Surname = updatePersonData.Surname
	person.Patronymic = updatePersonData.Patronymic
	person.Age = updatePersonData.Age
	person.Gender = updatePersonData.Gender
	person.Nationalize = updatePersonData.Nationalize

	db.Save(&person)

	return c.JSON(fiber.Map{"status": "success", "message": "Person Data Update", "data": person})
}

func (p Person) DeletePerson(c *fiber.Ctx) error {
	c.GetRespHeader("Access-Control-Allow-Origin", "*")
	db := database.DB
	var person = p.person

	id := c.Params("personId")

	db.Find(&person, "id = ?", id)

	if person.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No person present", "data": nil})
	}
	err := db.Delete(&person, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Delete Person"})
}

func (p Person) GetPerson(c *fiber.Ctx) error {
	c.GetRespHeader("Access-Control-Allow-Origin", "*")
	db := database.DB
	var person = p.person

	id := c.Params("personId")
	fmt.Println(id)
	db.Find(&person, "id = ?", id)

	if person.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No person present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Person Found", "data": person})
}

// Функция пагинации

func (p Person) Pagination(c *fiber.Ctx) error {
	c.GetRespHeader("Access-Control-Allow-Origin", "*")
	db := database.DB
	var persons = p.persons

	getPage := c.Params("page")
	page, err := strconv.ParseInt(getPage, 10, 32)
	realPage := int(page)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No №page for output", "data": nil})
	}
	getLimit := c.Params("limit")
	limit, err := strconv.ParseInt(getLimit, 10, 32)
	realLimit := int(limit)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No limit for output", "data": nil})
	}
	countPages := func() (int64, error) {
		var count int64
		db.Raw("SELECT COUNT(*) FROM Person_infos").Scan(&count)
		a := count / limit
		if (count % limit) == 0 {
			return a, nil
		} else {
			return a + 1, nil
		}
	}
	count, err := countPages()
	if err != nil {
		return c.JSON(fiber.Map{"status": "failed", "message": "Ashybka", "data": nil, "countPages": nil})
	}
	db.Limit(realLimit).Offset(realLimit * (realPage - 1)).Find(&persons)
	return c.JSON(fiber.Map{"status": "success", "message": "Person Found", "data": persons, "countPages": count})
}
