package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jimsyyap/blog/database"
)

func init() {
    database.ConnectDB()
}

func main() {
    sqlDB, err := database.DBConn.DB()

    if err != nil {
        panic(err)
    }

    defer sqlDB.Close()

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
