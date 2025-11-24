package main

import "fmt"

func main() {
	/*score :=0

	if score > 10 {
		fmt.Println("Ты красавчик")
	} else {
		fmt.Println("Поучись еще")
	}
	*/

	/*
		score := 10

		if score >= 10 {
			fmt.Println("Ты красавчик")
		} else {
			fmt.Println("Слабак")
		}
	*/

	/* sunny := false
	weekend := true

	if sunny && weekend {
		fmt.Println("Я иду гулять!")
	} else {
		fmt.Println("Я сопротивляюсь!")
	}

	ComputerClub := false
	icecream := true

	if ComputerClub || icecream {
		fmt.Println("Я иду гулять!")
	} else{
		fmt.Println("Я сопротивляюсь!")
	}
	*/

	//функция принимает на вход в параметрах строку(название месяца) и число (номер года)
	//и возвращает true если в переданном месяце 31 день и переданный год високосный, иначе false

	Mesyac := "Январь"
	God := 2000

	isVisokosny := God%4 == 0
	if isVisokosny == true && Mesyac == "Январь" || Mesyac == "Август" {
		fmt.Println("Все четко")
	} else {
		fmt.Println("Не четко")
	}
}
