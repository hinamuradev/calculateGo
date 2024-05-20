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

		var a, b, result int
		var operator string
		var isRoman bool

		// Используем регулярное выражение для разделения операндов и оператора
		roman := regexp.MustCompile(`^([IVXLCDM]+)([+\-*/])([IVXLCDM]+)$`)
		arabic := regexp.MustCompile(`^(\d+)([+\-*/])(\d+)$`)
		if roman.MatchString(text) {
			isRoman = true
			matches := roman.FindStringSubmatch(text)

			a = romanToArabic(matches[1])
			b = romanToArabic(matches[3])
			fmt.Println("Число 1.1:", a)
			fmt.Println("Оператор:", matches[2])
			fmt.Println("Число 2:", b)
			operator = matches[2]

		} else if arabic.MatchString(text) {
			matches := arabic.FindStringSubmatch(text)
			a, _ = strconv.Atoi(matches[1])
			b, _ = strconv.Atoi(matches[3])
			fmt.Println("Число 1.2:", a)
			fmt.Println("Оператор:", matches[2])
			fmt.Println("Число 2:", b)
			operator = matches[2]

		} else {
			fmt.Println("Некорректный ввод.")

		}

		fmt.Println("Число 1.3:", a)
		fmt.Println("Оператор:", operator)
		fmt.Println("Чиисло 2:", b)
		if a <= 10 && b <= 10 {
			switch operator {
			case "+":
				result = a + b
				fmt.Println("Твой ответ", result)
			case "-":
				result = a - b
				fmt.Println("Твой ответ", result)
			case "*":
				result = a * b
				fmt.Println("Твой ответ", result)
			case "/":
				if b != 0 {
					result = a / b
					fmt.Println("Твой ответ", result)
				} else {
					fmt.Println("На нуль делить нельзя")
					continue
				}
			default:
				fmt.Println("Ты допустил ошибку!")
			}

			if isRoman {
				fmt.Println("Ваше значение:", arabicToRoman(result))
			} else {
				fmt.Println("Ваше значение:", result)
			}
		} else {
			fmt.Println("Числа больше 10!")
		}
	}
}
