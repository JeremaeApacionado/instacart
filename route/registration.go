package route

import (
	"regexp"
	"instacart/database"
	"instacart/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CustomerReg(c *fiber.Ctx) error {
	var customer models.Customer
	new_customer := new(models.Customer)

	if err := c.BodyParser(&new_customer); err != nil {
		return c.Status(500).SendString("Server error")
	}
	regEmail := regexp.MustCompile("[a=zA-Z0-9_]+@[yahoogmail]+[.][com]{3}")
	formatterEmail := regEmail.MatchString(new_customer.Email)
	database.DB.Find(&customer, "email=?", new_customer.Email)
	database.DB.Find(&customer, "username=?", new_customer.Username)
	uniqueEmail := new_customer.Email != customer.Email
	uniqueUsername := new_customer.Username != customer.Username
	usernameLength := len(new_customer.Username) >= 8
	passwordLength := len(new_customer.Password) >= 8
	hash, _ := HashPasswordC(new_customer.Password)
	new_customer.Password = hash

	if  formatterEmail&& uniqueEmail && uniqueUsername && usernameLength && passwordLength {
		database.DB.Create(&new_customer)
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
		"CUSTOMER":    new_customer,
	})

}

func HashPasswordC(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GetCustomer(c *fiber.Ctx) error {
	var customer []models.Customer
	database.DB.Find(&customer)
	if len(customer) == 0 {
		return c.JSON(&fiber.Map{
			"Message": "User Does not Exist!",
		})
	}
	return c.JSON(&fiber.Map{
		"Users": customer,
	})
}
