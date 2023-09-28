package main

import "github.com/gofiber/fiber/v2/log"

func main() {
	app := AppFiber()

	log.Info("\nStarting server\nListen to the port 3000...")
	app.Listen(":3000")
}
