package boot

import (
	"log"
	"os"

	"github.com/RianIhsan/gofiber-relational-db/config"
	"github.com/RianIhsan/gofiber-relational-db/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func BootApp() {
	//LOAD ENV FILE
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}

	//CHECK DEFAULT PORT
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		config.PORT = portEnv
	}

	config.BootDatabase()
	config.ConnectDB()
	config.GoMigration()

	app := fiber.New()

	//Cors config
	/*app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
		ExposeHeaders:    config.ExposeHeaders,
		MaxAge:           config.MaxAge,
	}))*/

	//Route Init
	routes.SetupRoute(app)

	app.Listen(config.PORT)
}
