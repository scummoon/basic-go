package main

import "fmt"

func main() {
	year := 2026
	fmt.Println(year)
	if isLeapYear(year) {
		fmt.Println("Год високосный")
	} else {
		fmt.Println("Год не високосный")
	}
}

func isLeapYear(year int) bool {
	sum := year%4 == 0 && year%100 != 0 || year%100 == 0
	return sum

}
