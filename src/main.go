package main

import (
	bridge "fiber-queue/src/bridge"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	bridgeQueue := bridge.CreateBridge()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Running")
	})

	app.Get("/delay", func(c *fiber.Ctx) error {
		time.Sleep(300 * time.Millisecond)

		return c.SendString("Ok")
	})

	app.Get("/bridge", func(c *fiber.Ctx) error {
		response := make(chan string)

		bridgeQueue <- bridge.Message{
			Content:  "Hi",
			Response: response,
		}

		message := <-response

		return c.SendString(message)
	})

	app.Listen(":3000")
}
