package handler

import (
	"al-sufiaan-school-backend/apiType"
	"al-sufiaan-school-backend/database"
	"al-sufiaan-school-backend/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AddClass(c *fiber.Ctx) error {
	class := new(model.Class)
	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Unable to parse request body", "data": err})
	}
	class.Name = strings.ToUpper(class.Name)
	db := database.DB
	db.Create(&class)
	return c.JSON(fiber.Map{"status": "success", "message": "Class added successfully", "data": class})
}

func GetAllClass(c *fiber.Ctx) error {
	var classes []model.Class
	db := database.DB
	db.Where(&model.Class{SchoolId: 1}).Find(&classes)
	var data []apiType.ListClassResponse
	for _, class := range classes {
		var c apiType.ListClassResponse
		c.Name = class.Name
		c.SchoolId = class.SchoolId
		data = append(data, c)
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "List of classes",
		"data":    data,
	})

}
