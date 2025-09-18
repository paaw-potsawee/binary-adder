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
	Option string `json:"option"`
}

type CheckRequest struct {
	A string `json:"a"` 
	B string `json:"b"` 
	Answer string `json:"answer"` 
	Option string `json:"option"`
}

type CheckResponse struct {
	IsCorrect bool `json:"is_correct"`
}

func GetQuiz(c *fiber.Ctx) error {
	randombyte1 := rand.Intn(256)
	randombyte2 := rand.Intn(256)
	A := fmt.Sprintf("%08b",randombyte1)
	B := fmt.Sprintf("%08b",randombyte2)
	
	options := []string{"add","sub","xor","shift"}
	randomOption := options[rand.Intn(len(options))]
	quiz := Quiz{A: A, B: B,Option:randomOption}

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


	var hexStr string 
	switch req.Option {
	case "add":
		sum := (a + b) & 0xFF
		hexStr = fmt.Sprintf("%02x",sum)
	case "sub":
		result := (a - b) & 0xFF
		hexStr = fmt.Sprintf("%02x",result)
	case "xor":
		result := (a ^ b)
		hexStr = fmt.Sprintf("%02x",result)
	case "shift":
		result := (a << 1) & 0xFF
		hexStr = fmt.Sprintf("%02x",result)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid input option"})
	}



	res := CheckResponse{IsCorrect: hexStr == req.Answer}

	return c.JSON(res)
}
