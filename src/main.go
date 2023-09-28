package main

import "github.com/gofiber/fiber/v2/log"

func main() {
	OutputLog()
	app := AppFiber()

	// starting server
	log.Info("\nStarting server\nListen to the port 3000...")
	app.Listen(":3000")
}
