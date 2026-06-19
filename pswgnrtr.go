package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

type PasswordGenerator struct {
	Length     int
	UseSpecial bool
	Chars      string
}

func NewPasswordGenerator(length int, useSpecial bool) *PasswordGenerator {
	chars := lowercase + uppercase + digits
	if useSpecial {
		chars += special
	}
	return &PasswordGenerator{
		Length:     length,
		UseSpecial: useSpecial,
		Chars:      chars,
	}
}

func (pg *PasswordGenerator) getRandomChar() rune {
	index := rand.Intn(len(pg.Chars))
	return rune(pg.Chars[index])
}

func (pg *PasswordGenerator) Generate() string {
	password := make([]rune, pg.Length)
	for i := 0; i < pg.Length; i++ {
		password[i] = pg.getRandomChar()
	}
	return string(password)
}

func getArgValue(args []string, flag string) (string, bool) {
	for i, arg := range args {
		if arg == flag && i+1 < len(args) {
			return args[i+1], true
		}
	}
	return "", false
}

func parseLength(args []string) int {
	lengthStr, found := getArgValue(args, "--length")
	if !found {
		return 16
	}
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		fmt.Printf("Invalid length: %s\n", lengthStr)
		os.Exit(1)
	}
	return length
}

func parseSpecial(args []string) bool {
	for _, arg := range args {
		if arg == "--special" {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("Password Generator CLI")
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --length <number>  Length of password (default: 16)")
	fmt.Println("  --special          Include special characters")
}

func analyzePassword(password string) {
	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false
	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}
	fmt.Printf("Lowercase: %v | Uppercase: %v | Digits: %v | Special: %v\n",
		hasLower, hasUpper, hasDigit, hasSpecial)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	length := parseLength(os.Args)
	useSpecial := parseSpecial(os.Args)

	if length < 4 || length > 128 {
		fmt.Println("Password length must be between 4 and 128")
		os.Exit(1)
	}

	generator := NewPasswordGenerator(length, useSpecial)
	password := generator.Generate()

	fmt.Printf("\nGenerated Password (length %d):\n", len(password))
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(password)
	fmt.Println(strings.Repeat("-", 40))

	analyzePassword(password)
}
