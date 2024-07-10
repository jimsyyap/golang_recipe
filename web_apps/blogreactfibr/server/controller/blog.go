package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jimsyyap/blog/database"
    "github.com/jimsyyap/blog/model"
)

func BlogList (c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "Blog List",
        "msg": "Blog List",
    }

    //time.Sleep(time.Millisecond * 1000)

    db := database.DBConn
    var records []model.Blog
    db.Find(&records)
    context["blog_records"] = records

    c.Status(200)
    return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "Blog Create",
        "msg": "Blog Create",
    }
    c.Status(200)
    return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "Blog Update",
        "msg": "Blog Update",
    }
    c.Status(200)
    return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "Blog Delete",
        "msg": "Blog Delete",
    }
    c.Status(200)
    return c.JSON(context)
}


func BlogDetail(c *fiber.Ctx) error {
    context := fiber.Map{
        "statusText": "Blog Detail",
        "msg": "Blog Detail",
    }
    c.Status(200)
    return c.JSON(context)
}

/*
func BlogDetail(c *fiber.Ctx) error {
	// Implementation for getting a specific blog
	id := c.Params("id")
	return c.SendString("Blog detail for ID: " + id)
}
*/
