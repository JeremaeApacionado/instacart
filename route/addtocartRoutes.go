package route

import (
	"instacart/database"
	"instacart/model"

	"github.com/gofiber/fiber/v2"
)
//add to cart
func Getcart(c *fiber.Ctx) error {
	je := []model.Addtocart{}
	database.DB.Find(&je)
	return c.JSON(je)
}

func Getcarts(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	je := model.Addtocart{}
	database.DB.Find(&je, "atc_id", id)
	return c.JSON(je)

}
