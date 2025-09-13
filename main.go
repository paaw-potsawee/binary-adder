package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paaw-potsawee/binary-adder/internal/handlers/quiz"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web/static")

	app.Get("/quiz", quiz.GetQuiz)
	app.Post("/quiz/check", quiz.CheckQuiz)

	log.Println("Listening on :3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
    }
}
