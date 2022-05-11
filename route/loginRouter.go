package route

import (
	"fmt"
	"log"
	"instacart/database"
	"instacart/model"
	"instacart/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Log(c *fiber.Ctx) error {
	var log models.Login
	var user models.User
	util.BodyParser(c, &log)
	
	database.DB.Find(&user, "username=?", log.Username)
	if log.Username != user.Username {
		return c.JSON(&fiber.Map{
			"message":       "Wrong Username or Password",
			"login_success": false,
		})
	} else {
		match := CheckPasswordHash([]byte(user.Password), []byte(log.Password))
		if !match {
			return c.JSON(&fiber.Map{
				"message": "Wrong Username or Password",
				"success": false,
			})
		}
		fmt.Print("password match", match)
		response := CreateUserResponse(user)
		return c.JSON(&fiber.Map{

			"success": true,
			"message": "Login Success",
			"data":    response,
		})
	}

}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Println("Unable to compare password", err)
		return false
	}
	return true
}