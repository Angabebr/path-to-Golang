package main

import (
	"fmt"
)

func main() {
	text1 := "Get ready"

	text2 := "Game Over"

	score := 11

	ostatok := score % 3

	result := float64(score) / 3

	drob := 0.5

	answer := true

	fmt.Println(result)

	fmt.Println(drob)

	fmt.Println(ostatok)

	fmt.Println(answer)

	fmt.Println(text1)

	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1

	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1

	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы пролетели через трубу")
	score = score + 1

	fmt.Println("Ваш счет:", score)

	fmt.Println("Вы врезались в трубу\n", text2)

}
