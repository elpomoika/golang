package main

import "fmt"

func main() {
	first_num := 0
	fmt.Println("Введи первое число: ")
	fmt.Scan(&first_num)
	second_num := 0
	fmt.Println("Введи второе число: ")
	fmt.Scan(&second_num)

	sign := ""
	fmt.Println("Введи знак операции (-, +, :, *): ")
	fmt.Scan(&sign)

	var res float64

	if sign == "-" {
		res = float64(first_num) - float64(second_num)
		fmt.Println(res)
	} else if sign == "+" {
		res = float64(first_num) + float64(second_num)
		fmt.Println(res)
	} else if sign == ":" {
		res = float64(first_num) / float64(second_num)
		if second_num == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(res)
		}
	} else if sign == "*" {
		res = float64(first_num) * float64(second_num)
		fmt.Println(res)
	}

	fmt.Scan(&sign)
}
