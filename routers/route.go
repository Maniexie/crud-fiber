package routers

import (
	"github.com/Maniexie/crud-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/user", controllers.UserControllerShow)
	c.Post("/user", controllers.UserControllerCreate)
	c.Get("/user/:id", controllers.UserControllerFind)
	c.Put("/user/:id", controllers.UserControllerUpdate)
	c.Delete("/user/:id", controllers.UserControllerDelete)
}
