package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jimsyyap/blog/database"
    "github.com/jimsyyap/blog/router"
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

    router.SetupRoutes(app)
    app.Listen(":3000")
}
