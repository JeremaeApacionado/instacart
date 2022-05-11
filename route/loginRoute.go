package route

import(
	"instacart/database"
	"instacart/models"

	"github.com/gofiber/fiber/v2"
)

func Getinput(c *fiber.Ctx) error {
	var users []models.Login
	database.DB.Find(&users)
	if len(users) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User Does not Exist!",
		})
	}
	return c.JSON(&fiber.Map{
		"Users": users,
	})
}