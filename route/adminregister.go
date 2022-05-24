package route

import (
	"instacart/database"
	"instacart/models"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AdminReg(c *fiber.Ctx) error {
	var admin models.Admin
	new_admin := new(models.Admin)

	if err := c.BodyParser(&new_admin); err != nil {
		return c.Status(500).SendString("Server error")
	}
	regEmail := regexp.MustCompile("[a=zA-Z0-9_]+@[yahoogmail]+[.][com]{3}")
	formatterEmail := regEmail.MatchString(new_admin.Email)
	database.DB.Find(&admin, "email=?", new_admin.Email)
	database.DB.Find(&admin, "username=?", new_admin.Username)
	uniqueEmail := new_admin.Email != admin.Email
	uniqueUsername := new_admin.Username != admin.Username
	usernameLength := len(new_admin.Username) >= 8
	passwordLength := len(new_admin.Password) >= 8
	hash, _ := HashPasswordA(new_admin.Password)
	new_admin.Password = hash

	if  formatterEmail&& uniqueEmail && uniqueUsername && usernameLength && passwordLength {
		database.DB.Create(&new_admin)
	} else {
		if !uniqueEmail {
			return c.SendString("Email already exist!")
		}
		if !uniqueUsername {
			return c.SendString("Username already exist!")
		}
		if !usernameLength {
			return c.SendString("Username length should be atleast 8 characters!")
		}
		if !passwordLength {
			return c.SendString("Password length should be atleast 8 characters!")
		}
	}

	return c.JSON(&fiber.Map{
		"message": "User successfully registered as Customer",
		"CUSTOMER":    new_admin,
	})

}

func HashPasswordA(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GetAdmin(c *fiber.Ctx) error {
	var admin []models.Admin
	database.DB.Find(&admin)
	if len(admin) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User Does not Exist!",
		})
	}
	return c.JSON(&fiber.Map{
		"Admins": admin,
	})
}