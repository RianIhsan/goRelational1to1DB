package controllers

import (
	"github.com/RianIhsan/gofiber-relational-db/database"
	"github.com/RianIhsan/gofiber-relational-db/models"
	"github.com/gofiber/fiber/v2"
)

func LockerGetAll(c *fiber.Ctx) error {
  var locker []models.Locker

  database.DB.Preload("User").Find(&locker)

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message":"Locker tersedia",
    "locker": locker,
  })
}

func CreateLocker(c *fiber.Ctx) error {
  var locker models.Locker

  if err := c.BodyParser(&locker); err != nil {
    c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Gagal Parsing data JSON",
    })
  }

 if locker.Code == "" {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"Code is Required!",
    })
  }

 if locker.UserID == 0 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message":"USERID is Required",
    })
  }

  if err := database.DB.Create(&locker).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message":"Internal Server ERROR!",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "message":"Locker berhasil dibuat",
  })


}
