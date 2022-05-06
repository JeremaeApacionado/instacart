package route

import (
	"instacart/database"
	"instacart/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)
//add to cart
func Getcart(c *fiber.Ctx) error {
	je := []models.Addtocart{}
	database.DB.Find(&je)
	return c.JSON(je)
}

func Getcarts(c *fiber.Ctx) error {
	je := models.Addtocart{}
	user := models.User{}
	cartProduct := models.Product{}
	database.DB.Find(&je, "user_id", 2)
	database.DB.Find(&user, "user_id", je.UserID)
	database.DB.Find(&cartProduct, "product_id", je.ProductID)
	fmt.Println(je)
	fmt.Println(user)
	fmt.Println(cartProduct)

	return c.JSON(je)

}
