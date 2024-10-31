package main

import (
	"fmt"
)

func main() {
	списокПокупок := []string{}
	var вибір int

	for {
		fmt.Println("\n1. Додати товар")
		fmt.Println("2. Видалити товар")
		fmt.Println("3. Показати список")
		fmt.Println("4. Вийти")
		fmt.Print("Оберіть дію: ")
		fmt.Scan(&вибір)

		switch вибір {
		case 1:
			var товар string
			fmt.Print("Введіть назву товару: ")
			fmt.Scan(&товар)
			списокПокупок = append(списокПокупок, товар)
		case 2:
			var індекс int
			fmt.Print("Введіть індекс товару для видалення: ")
			fmt.Scan(&індекс)
			if індекс >= 0 && індекс < len(списокПокупок) {
				списокПокупок = append(списокПокупок[:індекс], списокПокупок[індекс+1:]...)
			} else {
				fmt.Println("Неправильний індекс")
			}
		case 3:
			fmt.Println("Список покупок:")
			for i, товар := range списокПокупок {
				fmt.Printf("%d: %s\n", i, товар)
			}
		case 4:
			return
		default:
			fmt.Println("Неправильний вибір")
		}
	}
}
