package quiz

import (
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
	mode := c.Query("mode","random")
	options := []string{"add","sub","xor","shift"}
	quiz,err := genQuiz(mode,&options) 
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error":"invalid mode"})	
	}
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

	res,err := checkQuiz(req.Option,a,b,req.Answer) 
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	return c.JSON(res)
}
