package route

import (
	"instacart/database"
	"instacart/models"

	"github.com/gofiber/fiber/v2"
)
//images
func GetImages(c *fiber.Ctx) error {
	je := []models.Images{}
	database.DB.Find(&je)
	return c.JSON(je)
}

func GetImage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	je := models.Images{}
	database.DB.Find(&je, "image_id", id)
	return c.JSON(je)

}

