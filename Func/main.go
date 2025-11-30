package main

import (
	"fmt"
)

func main() {
	var choice int
	var a, b int
	for {
		fmt.Println("===КАЛЬКУЛЯТОР===")
		fmt.Println("1. Сложение")
		fmt.Println("2. Вычитание")
		fmt.Println("3. Умножение")
		fmt.Println("4. Деление")
		fmt.Println("0. Выход")
		fmt.Println("Выберите номер операции: ")
		fmt.Scan(&choice)

		if choice == 0 {
			fmt.Println("Программа завершена")
			break
		}

		if choice < 1 || choice > 4 {
			fmt.Println("Неверный выбор! попробуйте заново")
			continue
		}
		fmt.Print("Введите первое число: ")
		fmt.Scan(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			plus(a, b)
		case 2:
			minus(a, b)
		case 3:
			multiply(a, b)
		case 4:
			divide(a, b)
		}
		fmt.Print("\nНажмите Enter для продолжения...")
		fmt.Scanln()
		fmt.Scanln()
	}
}

func plus(a, b int) {
	fmt.Println("===СЛОЖЕНИЕ===")
	fmt.Println("Твои входные данные= ", a, ",", b)
	fmt.Println("Сейчас посчитаем:", a+b)
}

func minus(a, b int) {
	fmt.Println("===ВЫЧИТАНИЕ===")
	fmt.Println("Твои входные данные= ", a, ",", b)
	fmt.Println("Сейчас посчитаем:", a-b)
}

func multiply(a, b int) {
	fmt.Println("===УМНОЖЕНИЕ===")
	fmt.Println("Твои входные данные= ", a, ",", b)
	fmt.Println("Сейчас посчитаем:", a*b)
}

func divide(a, b int) {
	fmt.Println("===ДЕЛЕНИЕ===")
	fmt.Println("Твои входные данные= ", a, ",", b)
	fmt.Println("Сейчас посчитаем:", a/b)
}
