// package main
// import("fmt")

// func main()  {
// 	fmt.Println("hello")
// }

package main

import (
	"github.com/Maniexie/crud-fiber/database"
	"github.com/Maniexie/crud-fiber/database/migration"
	"github.com/Maniexie/crud-fiber/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	migration.RunMigrate()
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "messasdsage",
	// 		"key":     "value",
	// 	})
	// })

	routers.RouterApp(app)

	app.Listen(":3000")
}
