package main

import (
	"fmt"
	"os"
	"strings"
)

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	arabic := 0
	prevValue := 0
	for _, char := range roman {
		value := romanNumerals[char]
		arabic += value
		if prevValue < value {
			arabic -= 2 * prevValue
		}
		prevValue = value
	}
	return arabic
}

func operate(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неподдерживаемая операция")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic("Неверное количество аргументов")
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		a = romanToArabic(args[0])
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		b = romanToArabic(args[2])
	}

	operator := args[1]
	if !strings.Contains("+-*/", operator) {
		panic("Неподдерживаемая операция")
	}

	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		panic("Числа должны быть от 1 до 10 включительно")
	}

	result := operate(operator, a, b)

	if result < 1 {
		panic("Результат не может быть меньше 1 для римских чисел")
	}

	fmt.Println(result)
}
