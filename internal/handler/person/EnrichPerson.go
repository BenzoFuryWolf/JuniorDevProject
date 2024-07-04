package personHandler

import (
	"encoding/json"
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/config"
	"github.com/gofiber/fiber/v2"
)

type API struct {
	sourse string
}

type PersonEnricher interface {
	nazy(name string) string
	myAge(name string) string
	myGender(name string) string
}

// Функция подключается к api и получает наиболее вероятную национальность по имени
func (a API) nazy(name string) string {
	a.sourse = config.Config("API_NAZI")
	url := fmt.Sprintf("%s%s", a.sourse, name)
	connect := fiber.Get(url)
	_, body, errs := connect.Bytes()
	if len(errs) > 0 {
		return errs[0].Error()
	}
	defer connect.ConnectionClose()
	var person map[string]interface{}

	err := json.Unmarshal(body, &person)
	if err != nil {
		return err.Error()
	}
	countries := person["country"].([]interface{})
	countryId := countries[0].(map[string]interface{})
	nazy := countryId["country_id"].(string)
	return nazy
}

// Функция подключается к api и получает наиболее вероятный возраст
func (a API) myAge(name string) string {
	a.sourse = config.Config("API_AGE")
	url := fmt.Sprintf("%s%s", a.sourse, name)
	connect := fiber.Get(url)
	_, body, errs := connect.Bytes()
	if len(errs) > 0 {
		return errs[0].Error()
	}
	defer connect.ConnectionClose()

	var person map[string]interface{}

	err := json.Unmarshal(body, &person)
	if err != nil {
		return err.Error()
	}

	age := person["age"].(float64)
	rAge := fmt.Sprintf("%.0f", age)
	return rAge
}

// Функция подключается к api и получает наиболее вероятный пол объекта
func (a API) myGender(name string) string {
	a.sourse = config.Config("API_GENDER")
	url := fmt.Sprintf("%s%s", a.sourse, name)
	connect := fiber.Get(url)
	_, body, errs := connect.Bytes()
	if len(errs) > 0 {
		return errs[0].Error()
	}
	defer connect.ConnectionClose()

	var person map[string]interface{}

	err := json.Unmarshal(body, &person)
	if err != nil {
		return err.Error()
	}
	gender := person["gender"].(string)
	return gender
}
