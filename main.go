package main

import (
	"github.com/gofiber/fiber/v2"
)

type Donor struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Dob     string `json:"dob"`
	Blood   string `json:"blood"`
	Disease bool   `json:"disease"`
	Rating  int    `json:"rating"`
}

var donors = []Donor{
	{Id: 1, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 2, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 3, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 4, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 5, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 6, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 7, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
	{Id: 8, Name: "name", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Rating: 4},
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/dev/donors", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(donors)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func NewDonor(id int, name string, phone string, dob string, blood string, disease bool, rating int) *Donor {
	return &Donor{Id: id, Name: name, Phone: phone, Dob: dob, Blood: blood, Disease: disease, Rating: rating}
}
