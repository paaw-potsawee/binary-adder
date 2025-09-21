package quiz

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func stringInSlice(a string, list *[]string) bool {
	return slices.Contains(*list, a)
}

func genQuiz(mode string, options *[]string) (*Quiz, error) {
	var option string
	randombyte1 := rand.Intn(256)
	randombyte2 := rand.Intn(256)
	A := fmt.Sprintf("%08b", randombyte1)
	B := fmt.Sprintf("%08b", randombyte2)
	if mode == "random" {
		idx := rand.Intn(len(*options))
		option = (*options)[idx]
	} else if stringInSlice(mode, options) {
		option = mode
	} else {
		return nil, errors.New("invalid mode")
	}
	quiz := Quiz{A: A, B: B, Option: option}
	return &quiz, nil
}

func checkQuiz(option string, a uint64, b uint64, answer string) (*CheckResponse, error) {
	var hexStr string
	switch option {
	case "add":
		sum := (a + b) & 0xFF
		hexStr = fmt.Sprintf("%02x", sum)
	case "sub":
		result := (a - b) & 0xFF
		hexStr = fmt.Sprintf("%02x", result)
	case "xor":
		result := (a ^ b)
		hexStr = fmt.Sprintf("%02x", result)
	case "shift":
		result := (a << 1) & 0xFF
		hexStr = fmt.Sprintf("%02x", result)
	default:
		return nil, errors.New("invalid input option")
	}
	return &CheckResponse{IsCorrect: hexStr == answer}, nil

}
