package route

import (
	"fmt"
	"instacart/database"

	"github.com/gofiber/fiber/v2"
)

func Order(c *fiber.Ctx) error {
	je := make(map[string]interface{})
	// database.DB.Debug().Select("user_id").Find(&je)
	database.DB.Debug().Table("products p").
		Select("p.product_id, p.name, p.description, p.images, p.stars, c.quantity_ordered").
		Joins("JOIN addtocarts c ON p.product_id = c.product_id").Find(&je)
	fmt.Println(je)
	return c.JSON(je)
}
