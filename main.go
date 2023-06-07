package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/sxc/cali/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	// app.Get("/foo", handleFoo)
	// rebuild the application
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUserByID)
	app.Listen(*listenAddr)
}

// func handleFoo(c *fiber.Ctx) error {
// 	return c.JSON(map[string]string{"msg": "working just fine!"})
// }

// func handleUser(c *fiber.Ctx) error {
// 	return c.JSON(map[string]string{"user": "John Appleseed"})
// }
