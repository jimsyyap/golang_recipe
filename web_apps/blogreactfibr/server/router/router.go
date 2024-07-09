package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jimsyyap/blog/controller"
)

func SetupRoutes(app *fiber.App) {
    /*
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    */
    app.Get("/", controller.BlogList)
    app.Get("/:id", controller.BlogDetail)
    app.Post("/", controller.BlogCreate)
    app.Put("/:id", controller.BlogUpdate)
    app.Delete("/:id", controller.BlogDelete)
}
