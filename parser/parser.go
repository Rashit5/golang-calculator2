package parser

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ParseExpression разбирает входное выражение на токены
func ParseExpression(input string) (string, string, string, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", "", "", errors.New("пустой ввод")
	}

	// Извлекаем первую строку
	firstStr, remainder, err := extractString(input)
	if err != nil {
		return "", "", "", err
	}

	// Извлекаем оператор
	op, remainder, err := extractOperator(remainder)
	if err != nil {
		return "", "", "", err
	}

	// Извлекаем второй аргумент (строку или число)
	secondArg, err := extractSecondArgument(remainder)
	if err != nil {
		return "", "", "", err
	}

	return firstStr, op, secondArg, nil
}

// extractString извлекает строку из ввода
func extractString(input string) (string, string, error) {
	if len(input) == 0 || input[0] != '"' {
		// Если первый символ не кавычка, вызываем панику
		panic("первый аргумент должен быть строкой")
	}

	input = input[1:] // Убираем открывающую кавычку
	endIndex := strings.Index(input, `"`)
	if endIndex == -1 {
		return "", "", errors.New("незакрытая кавычка")
	}

	str := input[:endIndex]
	if len(str) > 10 {
		return "", "", errors.New("длина строки не должна превышать 10 символов")
	}

	remainder := input[endIndex+1:]
	return str, remainder, nil
}

// extractOperator извлекает оператор
func extractOperator(input string) (string, string, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", "", errors.New("отсутствует оператор")
	}

	op := string(input[0])
	if op != "+" && op != "-" && op != "*" && op != "/" {
		return "", "", errors.New("неподдерживаемый оператор")
	}

	remainder := strings.TrimSpace(input[1:])
	return op, remainder, nil
}

// extractSecondArgument извлекает второй аргумент (строку или число)
func extractSecondArgument(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("отсутствует второй аргумент")
	}

	if input[0] == '"' {
		// Это строка
		str, _, err := extractString(input)
		return str, err
	} else if unicode.IsDigit(rune(input[0])) {
		// Это число
		numStr := ""
		for _, ch := range input {
			if unicode.IsDigit(ch) {
				numStr += string(ch)
			} else {
				break
			}
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", errors.New("некорректное число")
		}

		if num < 1 || num > 10 {
			return "", errors.New("число должно быть от 1 до 10")
		}

		return numStr, nil
	}

	return "", errors.New("некорректный второй аргумент")
}
