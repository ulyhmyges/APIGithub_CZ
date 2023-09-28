package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func AppFiber() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Get("/", handler)
	app.Get("/display/:user", handlerDisplay)

	return app
}

func handler(c *fiber.Ctx) error {
	fmt.Println("GET request / from: ", c.IP())
	return c.SendString("Welcome home ")
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
}

func handlerDisplay(c *fiber.Ctx) error {
	listRepos := DisplayRepos(Githubrepos(c.Params("user")))
	return c.Status(200).SendString(listRepos)
}
