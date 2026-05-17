package main

import "fmt"

func main() {
	greeting("olzhas", 24)
	add(2, 2)
	chet(24)
	umnozh(10)
	temp(15)
	checkPassword("12345")
	checkNumber(-5)
	checkage(19)
	checkDate(29, 02, 2026)
	calculator(20, 5, "*")
	checkGrade(77)
	multiplyTable(5)

	balance := 1000
	balance = withdraw(balance, 899)
	fmt.Println("Ваш баланс:", balance)

}

func greeting(name string, age int) {
	fmt.Println("Привет", name)
	fmt.Println("Мне", age, "года")
}

func add(sum int, sum2 int) {
	fmt.Println(sum + sum2)
}

func chet(num int) {
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}
func umnozh(number int) {
	fmt.Println(number * 10)
}

func temp(num1 int) {
	if num1 >= 25 {
		fmt.Println("Очень жарко")
	} else if num1 > 20 {
		fmt.Println("Тепло")
	} else if num1 > 10 {
		fmt.Println("холодно")
	} else {
		fmt.Println("Очень холодно")
	}
}

func checkPassword(password string) {
	realpassword := "12345"
	if password == realpassword {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}
}

func checkNumber(number int) {
	if number > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}

func checkage(age int) {
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}
}

func isLeapYear(year int) bool {
	sum := (year%4 == 0 && year%100 != 0) || year%400 == 0

	return sum
}

func checkDate(day int, month int, year int) {
	if year <= 0 {
		fmt.Println("invalid year value")
	} else if month < 1 || month > 12 {
		fmt.Println("invalid month value")
	} else if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) && (day > 31 || day < 1) {
		fmt.Println("Invalid date 31")
	} else if (month == 4 || month == 6 || month == 9 || month == 11) && (day > 30 || day < 1) {
		fmt.Println("Invalid date 30")
	} else if (month == 2) && isLeapYear(year) && (day > 29 || day < 1) {
		fmt.Println("Invalid date 29")
	} else if (month == 2) && !isLeapYear(year) && (day > 28 || day < 1) {
		fmt.Println("Invalid date 28")
	} else {
		fmt.Println("Correct")
	}
}

func calculator(a int, b int, op string) {
	if op == "+" {
		fmt.Println(a + b)
	} else if op == "-" {
		fmt.Println(a - b)
	} else if op == "*" {
		fmt.Println(a * b)
	} else if op == "/" {
		if b == 0 {
			fmt.Println("cannot divide by zero")
		} else {
			fmt.Println(a / b)
		}
	} else {
		fmt.Println("err")
	}
}

func checkGrade(score int) {
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score <= 89 {
		fmt.Println("B")
	} else if score >= 70 && score <= 79 {
		fmt.Println("C")
	} else if score >= 60 && score <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

func multiplyTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n * i)
	}
}

func withdraw(balance int, amount int) int {
	if amount <= 0 {
		fmt.Println("Invalid amount")
		return balance
	}
	if amount > balance {
		fmt.Println("Not enough money")
		return balance
	}
	fmt.Println("Вы успешно сняли деньги")
	return balance - amount

}
