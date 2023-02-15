package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mastama/go-patungan/config"
)

func main() {
	// call ConnectToDB
	db.ConnectToDB()

	app := fiber.New()

	app.Listen(":8081")
}
