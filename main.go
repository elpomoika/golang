package main

import (
	"fmt"
	"math"
)

func main() {
	first_num := 0
	fmt.Println("Введи первое число: ")
	fmt.Scan(&first_num)
	second_num := 0
	fmt.Println("Введи второе число: ")
	fmt.Scan(&second_num)

	sign := ""
	fmt.Println("Введи знак операции (-, +, /, *, **): ")
	fmt.Scan(&sign)

	var res float64
	var res1 float64

	switch sign {
	case "+":
		res = float64(first_num) + float64(second_num)
	case "-":
		res = float64(first_num) - float64(second_num) 
	case "/":
		res = float64(first_num) / float64(second_num) 
	case "*":
		res = float64(first_num) * float64(second_num) 
	case "**":
		degree := 1
		fmt.Println("В какую степень возводим: ")
		fmt.Scan(&degree)
		res = math.Pow(float64(first_num), float64(degree))
		res1 = math.Pow(float64(second_num), float64(degree))
		fmt.Println("Первое число в степени " + fmt.Sprint(degree) + ": " + fmt.Sprint(res))
		fmt.Println("Второе число в степени " + fmt.Sprint(degree) + ": " + fmt.Sprint(res1) )
	}
		

	fmt.Scan(&sign)
}
