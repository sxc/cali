package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sxc/cali/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName:  "Appleseed",
	}
	return c.JSON(u)
}

func HandleGetUserByID(c *fiber.Ctx) error {
	return c.JSON("John Appleseed")
}
