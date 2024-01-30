package personHandler

import (
	"encoding/json"
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/database"
	"github.com/BenzoFuryWolf/MyProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
)

func GetPersons(c *fiber.Ctx) error {
	var db = database.DB
	var person []model.Person_info

	db.Find(&person)

	if len(person) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No persons present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Persons found", "data": person})
}

func CreatePersons(c *fiber.Ctx) error {
	var db = database.DB
	var person = new(model.Person_info)
	err := c.BodyParser(person)
	//Функция подключается к api и получает наиболее вероятную национальность по имени
	nazy := func(string2 string) string {
		url := fmt.Sprintf("https://api.nationalize.io/?name=%s", string2)
		connect := fiber.Get(url)
		_, body, errs := connect.Bytes()
		if len(errs) > 0 {
			return errs[0].Error()
		}
		defer connect.ConnectionClose()
		var person map[string]interface{}

		err = json.Unmarshal(body, &person)
		if err != nil {
			return err.Error()
		}
		countries := person["country"].([]interface{})
		countryId := countries[0].(map[string]interface{})
		nazy := countryId["country_id"].(string)
		return nazy
	}
	//Функция подключается к api и получает наиболее вероятный возраст
	myAge := func(str string) string {
		url := fmt.Sprintf("https://api.agify.io/?name=%s", str)
		connect := fiber.Get(url)
		_, body, errs := connect.Bytes()
		if len(errs) > 0 {
			return errs[0].Error()
		}
		defer connect.ConnectionClose()

		var person map[string]interface{}

		err = json.Unmarshal(body, &person)
		if err != nil {
			return err.Error()
		}

		age := person["age"].(float64)
		rAge := fmt.Sprintf("%.0f", age)
		return rAge
	}
	//Функция подключается к api и получает наиболее вероятный пол объекта
	myGender := func(str string) string {
		url := fmt.Sprintf("https://api.genderize.io/?name=%s", str)
		connect := fiber.Get(url)
		_, body, errs := connect.Bytes()
		if len(errs) > 0 {
			return errs[0].Error()
		}
		defer connect.ConnectionClose()

		var person map[string]interface{}

		err = json.Unmarshal(body, &person)
		if err != nil {
			return err.Error()
		}
		gender := person["gender"].(string)
		return gender
	}
	nazionalize := nazy(person.Name)
	age, err := strconv.Atoi(myAge(person.Name))
	if err != nil {
		return err
	}
	gender := myGender(person.Name)
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

func UpdatePerson(c *fiber.Ctx) error {
	type updatePerson struct {
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
		Nationalize string `json:"nationalize"`
	}
	db := database.DB
	var person model.Person_info

	id := c.Params("personId")
	db.Find(&person, "id = ?", id)

	if person.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No person present", "data": nil})
	}

	var updatePersonData updatePerson
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

func DeletePerson(c *fiber.Ctx) error {
	db := database.DB
	var person model.Person_info

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
func GetPerson(c *fiber.Ctx) error {
	db := database.DB
	var person model.Person_info

	id := c.Params("personId")
	db.Find(&person, "id = ?", id)

	if person.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No person present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Person Found", "data": person})
}

// Функция пагинации
