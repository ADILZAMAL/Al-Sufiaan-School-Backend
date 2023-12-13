package handler

import (
	"al-sufiaan-school-backend/apiType"
	"al-sufiaan-school-backend/database"
	"al-sufiaan-school-backend/model"

	"github.com/gofiber/fiber/v2"
)

func OnboardSchool(c *fiber.Ctx) error {
	school := new(model.School)
	if err := c.BodyParser(&school); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on school onboard request", "data": err})
	}
	school.Active = true
	db := database.DB
	db.Create(&school)
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": school})
	// type LoginInput struct {
	// 	Identity string `json:"identity"`
	// 	Password string `json:"password"`
	// }
	// type UserData struct {
	// 	ID       uint   `json:"id"`
	// 	Username string `json:"username"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }
	// input := new(LoginInput)
	// var userData UserData

	// if err := c.BodyParser(&input); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	// }
	// identity := input.Identity
	// pass := input.Password
	// userModel, err := new(model.User), *new(error)
	// if isEmail(identity) {
	// 	userModel, err = getUserByEmail(identity)
	// } else {
	// 	userModel, err = getUserByUsername(identity)
	// }

	// if userModel == nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	// } else {
	// 	userData = UserData{
	// 		ID:       11231,
	// 		Username: userModel.Username,
	// 		Email:    userModel.Email,
	// 		Password: userModel.Password,
	// 	}
	// }

	// if !CheckPasswordHash(pass, userData.Password) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	// }

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = userData.Username
	// claims["user_id"] = userData.ID
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte(config.Config("SECRET")))
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

}

func GetSchool(c *fiber.Ctx) error {
	school := new(model.School)
	id := c.Params("id")
	db := database.DB
	db.Find(&school, id)

	var response apiType.GetSchoolResponse
	response.Id = school.Id
	response.Name = school.Name
	response.Address = school.Address
	response.Mobile = school.Mobile
	response.UdiceNo = school.UdiceNo
	response.Email = school.Email
	response.ActiveSession = school.ActiveSession
	return c.JSON(fiber.Map{"status": "success", "message": "School found", "data": response})
}
