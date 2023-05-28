package controllers

import (
	"github.com/RianIhsan/gofiber-relational-db/database"
	"github.com/RianIhsan/gofiber-relational-db/models"
	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx) error {
  var users []models.User

  database.DB.Preload("Locker").Find(&users)

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message":"Data tersedia di database",
    "data": users,
  })
}

func CreateUser(c *fiber.Ctx) error {
  user := new(models.User)

  if err := c.BodyParser(&user); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Gagal Parsing data JSON",
    })
  }

  if user.Name == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Name is Required!",
    })
  }

  if err := database.DB.Create(&user).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message":"Status internal Server ERROR!",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message":"User berhasil dibuat",
    "data":user,
  })
}
