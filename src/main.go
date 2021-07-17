package main

import (
	bridge "fiber-queue/src/bridge"

	"github.com/gofiber/fiber/v2"
)

func main() {
	bridgeQueue := bridge.CreateBridge()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Running")
	})

	app.Get("/bridge", func(c *fiber.Ctx) error {
		response := make(chan string)

		bridgeQueue <- bridge.BridgeMessage{
			Content:  "Hi",
			Response: response,
		}

		message := <-response

		return c.SendString(message)
	})

	app.Listen(":3000")
}
