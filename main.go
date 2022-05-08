package main

import (
	"fmt"
	"instacart/database"
	"instacart/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}

func Routes(app *fiber.App) {

	//users
	app.Post("/user", route.Registration)
	app.Get("/user", route.GetUsers)
	//Product
	app.Post("/product", route.AddProduct)
	app.Get("/product", route.GetProductName)
	app.Get("/product/:id", route.GetProduct)
	app.Delete("/product/:id", route.Delete)
	app.Put("/product/:id", route.Update)
	//addtocart
	app.Post("/addtocart", route.Getcart)
	app.Get("/addtocart/:id", route.Getcarts)
	app.Get("/sample/:id", route.Sample)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Could not load .env file")
	}
	port := os.Getenv("PORT")
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":" + PORT))
}
