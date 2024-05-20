package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func arabicToRoman(result int) string {
	val := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}

	var roman string
	for i := len(val) - 1; i >= 0; i-- {
		for result >= val[i] {
			result -= val[i]
			roman += symbols[i]
		}
	}
	return roman
}

func romanToArabic(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	total := 0
	prev := 0
	for _, char := range s {
		curr := romanMap[char]
		if curr > prev {
			total += curr - 2*prev
		} else {
			total += curr
		}
		prev = curr
	}
	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите выражение: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")

		// Используем регулярное выражение для разделения операндов и оператора
		re := regexp.MustCompile(`(\d+)([+\-*/])(\d+)`)
		matches := re.FindStringSubmatch(text)

		operand1, _ := strconv.Atoi(matches[1])
		operand2, _ := strconv.Atoi(matches[3])

		fmt.Println("Число 1:", operand1)
		fmt.Println("Оператор:", matches[2])
		fmt.Println("Чиисло 2:", operand2)
		if operand1 <= 10 && operand2 <= 10 {
			switch matches[2] {
			case "+":
				fmt.Println("Ваше значение: ", operand1+operand2)
				break
			case "-":
				fmt.Println("Ваше значение: ", operand1-operand2)
			case "*":
				fmt.Println("Ваше значение: ", operand1*operand2)
			case "/":
				if operand2 != 0 {
					fmt.Println("Ваше значение: ", operand1/operand2)
				} else {
					fmt.Println("На нуль делить нельзя")
				}
			default:
				fmt.Println("Ты допустил ошибку!")
			}
		} else {
			fmt.Println("Числа больше 10!")
		}

	}
}
