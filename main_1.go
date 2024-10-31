package main

import (
	"fmt"
)

// Функція для перевертання цифр числа
func перевернутиЧисло(число int) int {
	перевернуте := 0
	for число != 0 {
		перевернуте = перевернуте*10 + число%10
		число /= 10
	}
	return перевернуте
}

func main() {
	var число int
	fmt.Print("Введіть число: ")
	fmt.Scan(&число)
	fmt.Println("Перевернуте число:", перевернутиЧисло(число))
}
