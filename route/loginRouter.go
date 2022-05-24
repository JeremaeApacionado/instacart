package route

import (
	"fmt"
	"log"
	"instacart/database"
	"instacart/models"
	"instacart/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func Log(c *fiber.Ctx) error {
	var log Login
	var user models.Customer
	var user1 models.Admin
	util.BodyParser(c, &log)
	
	database.DB.Find(&user,&user1, "Email=?", log.Email)
	if log.Email != user.Email {
		return c.JSON(&fiber.Map{
			"message":       "Wrong Email or Password",
			"login_success": false,
		})
	}else if (user.Email == user1.Email){
		return c.JSON(&fiber.Map{
			"message": "Wrong Email or Password",
			"login_success": false,
		})
	}else {
		match := CheckPasswordHash([]byte(user.Password), []byte(log.Password))
		if !match {
			return c.JSON(&fiber.Map{
				"message": "Wrong Username or Password",
				"success": false,
			})
		}
		fmt.Print("password match", match)
		return c.JSON(&fiber.Map{

			"success": true,
			"message": "Login Success",
			"data":    user,
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