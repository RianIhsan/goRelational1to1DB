package routes

import (
	"github.com/RianIhsan/gofiber-relational-db/controllers"
	"github.com/gofiber/fiber/v2"
)

func allroutes(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
  app.Post("/users", controllers.CreateUser)

  app.Get("/lockers", controllers.LockerGetAll)
  app.Post("/lockers", controllers.CreateLocker)
}
