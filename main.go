package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/paaw-potsawee/binary-adder/internal/handlers/quiz"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web/static")

	app.Get("/quiz", quiz.GetQuiz)
	app.Post("/quiz/check", quiz.CheckQuiz)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
