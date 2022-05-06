package route

import (
	"instacart/database"
	"instacart/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)
//add to cart
func Getorder(c *fiber.Ctx) error {
	cart := []models.Cart{}
	database.DB.Find(&cart)
	return c.JSON(cart)
}

func Getorders(c *fiber.Ctx) error {
	cart := models.Cart{}
	database.DB.Find(&cart, "cart_id", )
	fmt.Println(cart)

	return c.JSON(cart)

}
