package main

import (
	"fmt"
	"os"
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
	app.Get("/login", route.Getinput)
	//Product
	app.Post("/product", route.AddProduct)
	app.Get("/product", route.GetProductName)
	app.Get("/product/:id", route.GetProduct)
	app.Delete("/product/:id", route.Delete)
	app.Put("/product/:id", route.Update)
	//addtocart
	app.Post("/addtocart", route.Getcart)
	app.Get("/addtocart/:id", route.Getcarts)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Could not load .env file")
	}
	PORT := os.Getenv("PORT")
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":" + PORT))
}
