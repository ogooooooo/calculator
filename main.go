package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidExpression(s string) bool { // Регулярное выражение для проверки формата строки
	re := regexp.MustCompile(`^"(.*?)"\s*([\+\-\*\/])\s*("(.*?)"|([1-9]|10))$`)
	return re.MatchString(s)
}

func splitted(str string) (string, string, string) { //разрезаем строку
	re := regexp.MustCompile(`(".*?"|\d+)\s*([-+*/])\s*(".*?"|\d+)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) == 4 {
		return matches[1], matches[2], matches[3]
	}

	return "", "", ""
}

func addition(st1, st2 string) string { // соединение(сложение) двух строк
	return st1[0:len(st1)-1] + st2[1:]
}

func division(str1, str2 string) string { // Деление строки на число
	number_str2, _ := strconv.Atoi(str2) // Преобразуем строку в целое число
	len_str1 := len(str1) - 2            //находим длину строки без учета двойных кавычек
	return str1[1 : len_str1/number_str2+1]
}

func multiplication(str1, str2 string) string { //умножение
	number_str2, _ := strconv.Atoi(str2)                    // Преобразуем строку в целое число
	str := strings.Repeat(str1[1:len(str1)-1], number_str2) // умножаем строку на число
	if len(str) > 40 {                                      // если строка больше 40 символов добавляем ...
		return str[:40] + "..."
	}
	return str
}

func main() {
	//str := `"Golang" * 5`

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Считываем строку
	str := scanner.Text()

	proverka := isValidExpression(str)
	if !proverka {
		panic("Некорректная строка")
	}

	str1, operator, str2 := splitted(str)

	if len(str1) > 12 || len(str2) > 12 {
		panic("Одна из строк больше 10 символов")
	}

	if operator == "+" { //сложение строк
		addition2 := addition(str1, str2)
		fmt.Println(addition2)
	} else if operator == "/" { //деление
		division2 := division(str1, str2)
		fmt.Println(`"` + division2 + `"`)
	} else if operator == "*" { // умножение
		multiplication := multiplication(str1, str2)
		fmt.Println(`"` + multiplication + `"`)

	} else if operator == "-" { //вычитание
		replacedText := strings.ReplaceAll(str1, str2[1:len(str2)-1], "")
		fmt.Println(replacedText)
	}

}
