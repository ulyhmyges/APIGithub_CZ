package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func fiberApp() {
	app := fiber.New(fiber.Config{})

	app.Get("/", handler)

}

func handler(c *fiber.Ctx) error {
	fmt.Println("GET request / from: ", c.IP())
	return c.SendString("GET / ")
}

func OutputLog() {
	// Output to ./test.log file
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error!")
		log.Fatal(err)
	}
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	//log.Info("info")
	//log.Warn("warn")
}
