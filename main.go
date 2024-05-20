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

		var (
			a, b, result int
			operator     string
			isRoman      bool
		)

		// Используем регулярное выражение для разделения операндов и оператора
		roman := regexp.MustCompile(`^([IVXLCDM]+)([+\-*/])([IVXLCDM]+)$`)
		arabic := regexp.MustCompile(`^(\d+)([+\-*/])(\d+)$`)
		if roman.MatchString(text) { // Римские числа
			isRoman = true
			matches := roman.FindStringSubmatch(text)

			a = romanToArabic(matches[1])
			b = romanToArabic(matches[3])
			operator = matches[2]

		} else if arabic.MatchString(text) { // Арабские числа
			matches := arabic.FindStringSubmatch(text)
			a, _ = strconv.Atoi(matches[1])
			b, _ = strconv.Atoi(matches[3])
			operator = matches[2]

		} else {
			panic("Некорректный ввод.")
		}

		if a <= 10 && b <= 10 {
			switch operator {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b != 0 {
					result = a / b
				} else {
					panic("На нуль делить нельзя")
				}
			default:
				fmt.Println("Ты допустил ошибку!")
				panic("Ты допустил ошибку!")
			}

			if isRoman {
				if result < 1 {
					panic("Римские число не может быть равно 0 или меньше его!!!")
				}
				fmt.Println("Ваше значение:", arabicToRoman(result))
			} else {
				fmt.Println("Ваше значение:", result)
			}
		} else {
			panic("Число больше 10!")
		}
	}
}
