package quiz

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Quiz struct {
	A string `json:"a"`
	B string `json:"b"`
}

type CheckRequest struct {
	A string `json:"a"` 
	B string `json:"b"` 
	Answer string `json:"answer"` 
}

type CheckResponse struct {
	IsCorrect bool `json:"is_correct"`
}

func GetQuiz(c *fiber.Ctx) error {
	randombyte1 := rand.Intn(256)
	randombyte2 := rand.Intn(256)
	A := fmt.Sprintf("%08b",randombyte1)
	B := fmt.Sprintf("%08b",randombyte2)
	
	quiz := Quiz{A: A, B: B}

	return c.JSON(quiz)
}

func CheckQuiz(c *fiber.Ctx) error {
	req := new(CheckRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	a,err := strconv.ParseUint(req.A,2,8)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid input"})
	}
	b,err := strconv.ParseUint(req.B,2,8)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid input"})
	}

	sum := (a + b) & 0xFF

	hexStr := fmt.Sprintf("%02x",sum)

	res := CheckResponse{IsCorrect: hexStr == req.Answer}

	return c.JSON(res)
}
