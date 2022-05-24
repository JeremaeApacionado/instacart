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

type ALogin struct {
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func ALog(c *fiber.Ctx) error {
	var Alog ALogin
	var admin models.Admin
	util.BodyParser(c, &Alog)
	
	database.DB.Find(&admin, "Email=?", Alog.Email)
	if Alog.Email != admin.Email {
		return c.JSON(&fiber.Map{
			"message":       "Wrong Email or Password",
			"login_success": false,
		})
	}else {
		match := CheckPasswordHash([]byte(admin.Password), []byte(Alog.Password))
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
			"data":    admin,
		})
	}

}
func CheckAdminPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		log.Println("Unable to compare password", err)
		return false
	}
	return true
}