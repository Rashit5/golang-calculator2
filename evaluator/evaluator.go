package evaluator

import (
	"errors"
	"strconv"
	"strings"
)

// Evaluate вычисляет результат выражения
func Evaluate(firstStr, op, secondArg string) (string, error) {
	switch op {
	case "+":
		// Сложение строк
		result := firstStr + secondArg
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		return result, nil
	case "-":
		// Вычитание строки из строки
		result := strings.ReplaceAll(firstStr, secondArg, "")
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		return result, nil
	case "*":
		// Умножение строки на число
		n, err := strconv.Atoi(secondArg)
		if err != nil {
			return "", errors.New("некорректное число")
		}
		result := strings.Repeat(firstStr, n)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		return result, nil
	case "/":
		// Деление строки на число
		n, err := strconv.Atoi(secondArg)
		if err != nil {
			return "", errors.New("некорректное число")
		}
		if n == 0 {
			return "", errors.New("деление на ноль")
		}
		result := firstStr[:len(firstStr)/n]
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		return result, nil
	default:
		return "", errors.New("неподдерживаемая операция")
	}
}
