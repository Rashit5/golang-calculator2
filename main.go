package main

import (
	"bufio"
	"fmt"
	"golang-calculator2/evaluator"
	"golang-calculator2/parser"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение (например, \"Hello\" + \"World\"): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Парсим выражение
	firstStr, op, secondArg, err := parser.ParseExpression(input)
	if err != nil {
		panic("Ошибка: " + err.Error())
	}

	// Вычисляем результат
	result, err := evaluator.Evaluate(firstStr, op, secondArg)
	if err != nil {
		panic("Ошибка: " + err.Error())
	}

	// Ограничиваем длину результата до 40 символов
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	// Выводим результат
	fmt.Printf("\"%s\"\n", result)
}
