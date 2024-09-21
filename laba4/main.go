package main

import (
	"fmt"
	"strings"
)

func main() {
	//1 номер
	peopleMap := map[string]int{
		"Александр Правин": 26,
		"Арбуз Арбуз":      1,
		"Денис Голиченко":  20,
		"Шэдоу Финд":       666,
	}

	fmt.Println(peopleMap)
	addPerson(&peopleMap, "Бородин Никита", 19)

	//2 номер
	fmt.Println(peopleMap, peopleMap["Арбуз Арбуз"])
	avgAge(peopleMap)

	//3 номер
	deleteFromMapByName(peopleMap, "Александр Правин")
	fmt.Println(peopleMap)

	//4 номер
	fmt.Println("4 Номер")
	var str string
	fmt.Scan(&str)
	makeUppercase(str)

	//5 номер
	fmt.Println("5 Номер")
	//Для окончания ввода чисел нужно написать 0
	var nums []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums = append(nums, input)
	}

	megaSum(nums)

	//6 номер
	//Для окончания ввода чисел нужно написать 0
	fmt.Println("6 номер")
	var nums2 []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		nums2 = append(nums2, input)
	}
	reverseArrOfNums(nums2)
}

func addPerson(peopleMap *map[string]int, name string, age int) {
	(*peopleMap)[name] = age
	fmt.Printf("Добавлен: %s, Возраст: %d\n", name, age)
}

func reverseArrOfNums(nums []int) {
	for _, v := range nums {
		defer fmt.Println(v)
	}
}

func megaSum(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}

func makeUppercase(str string) {
	fmt.Println(strings.ToUpper(str))
}

func deleteFromMapByName(mapExample map[string]int, str string) {
	fmt.Println("Before: ", mapExample)
	delete(mapExample, str)
	fmt.Println("After: ", mapExample)
}

func avgAge(mapExample map[string]int) {
	ages := 0
	counter := 0
	for _, v := range mapExample {
		ages += v
		counter++
	}
	fmt.Println(float64(ages) / float64(counter))
}
