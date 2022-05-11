package route

import(
"fmt"
"instacart/database"
"instacart/models"

"github.com/gofiber/fiber/v2"
"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user []models.User
	fmt.Println("Login data: ", data)
	database.Db.Where("login_username = ?", data["login_username"]).First(&user)
	fmt.Println("user:", user)
	if len(user) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(data["login_password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"Message": "user not found",
			"Error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login"})
}