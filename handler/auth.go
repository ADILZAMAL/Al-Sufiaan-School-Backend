package handler

import (
	"al-sufiaan-school-backend/config"
	"al-sufiaan-school-backend/database"
	"al-sufiaan-school-backend/model"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(email string) (*model.User, error) {
	db := database.DB
	var user model.User
	result := db.Where(&model.User{Email: email}).Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// func getUserByUsername(u string) (*model.User, error) {
// 	db := database.DB
// 	var user model.User
// 	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func isEmail(email string) bool {
// 	_, err := mail.ParseAddress(email)
// 	return err == nil
// }

// // Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// type UserData struct {
	// 	ID       uint   `json:"id"`
	// 	Username string `json:"username"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }
	input := new(LoginInput)
	// var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	email := input.Email
	pass := input.Password
	userModel, err := new(model.User), *new(error)
	userModel, err = getUserByEmail(email)
	fmt.Println(userModel)
	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if !CheckPasswordHash(pass, userModel.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userModel.Email
	claims["id"] = userModel.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":      "success",
		"message":     "Success login",
		"token":       t,
		"name":        userModel.Name,
		"email":       userModel.Email,
		"designation": userModel.Designation,
		"school_id":   userModel.SchoolId,
	})
}
