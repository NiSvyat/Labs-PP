package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//1 номер

	//Текущая дата
	fmt.Println("Текущая дата:")
	fmt.Println()

	//Текущее время
	fmt.Println("Текущее время:")
	fmt.Println(time.Time.Clock(time.Now()))

	//Дата и время
	fmt.Println("Текущие дата и время:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))

	//Номера 2 и 3
	intExample, float64Example, stringExample, boolExample := 1, 4.88, "ZV", false
	fmt.Println(intExample, float64Example, stringExample, boolExample)

	//Номера 4 и 5
	val1, _ := choseOperation(3, 2, "+")
	fmt.Println(val1)

	val2, _ := choseOperation(3, 2, "-")
	fmt.Println(val2)

	val3, _ := choseOperation(3, 2, "*")
	fmt.Println(val3)

	val4, err := choseOperation(3, 2, "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val4)

	//Проверка на ноль
	val5, err := choseOperation(3, 0, "/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val5)

	//6 номер

	var num1 int
	fmt.Scanln(&num1)

	var num2 int
	fmt.Scanln(&num2)

	var num3 int
	fmt.Scanln(&num3)

	fmt.Println(float64(num1+num2+num3) / 3)
}

func choseOperation(num1, num2 int, opertaion string) (float64, error) {
	switch opertaion {
	case "-":
		return float64(num1 - num2), nil
	case "+":
		return float64(num1 + num2), nil
	case "*":
		return float64(num1 * num2), nil
	case "/":
		if num2 != 0 {
			return float64(num1) / float64(num2), nil
		}
		return -1, errors.New("нельзя делить на 0, псих")
	}

	return -1, errors.New("ошибся, но где?...")
}
