package main

import "fmt"

func main() {
	color := "red"

	if color == "red" {
		fmt.Println("Стоп")
	} else if color == "yellow" {
		fmt.Println("Приготовься")
	} else if color == "green" {
		fmt.Println("Езжай")
	} else {
		fmt.Println("Неверный сигнал")
	}

	for i := 10; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Четное")
		} else {
			fmt.Println(i, "Нечетное")
		}
	}
	temp := 11

	if temp <= 10 {
		fmt.Println("Холодно")
	} else {
		fmt.Println("Тепло")
	}
}
