package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//1 номер
	evenCheck(1)
	evenCheck(2)

	//2 номер
	positiveCheck(-1)
	positiveCheck(2)
	positiveCheck(0)

	//3 номер
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//4 номер
	getLength("Shadow Fiend")
	getLength("Фишман, встань мид и стой со мной!")

	//5 номер
	rect1 := Rectangle{5, 2}
	rect1.GetRectangleArea()

	//6 номер
	avgNums(2, 3)
	avgNums(5, 6, 7)
}

func avgNums(num ...int) {
	sum := 0
	counter := 0
	for _, v := range num {
		counter++
		sum += v
	}
	fmt.Println(float64(sum) / float64(counter))
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) GetRectangleArea() {
	fmt.Println((float64(r.width) * float64(r.height)))
}

// Русские символы считаются за 2 при обычном len
func getLength(str string) {
	fmt.Println(utf8.RuneCountInString(str))
}

func evenCheck(num int) {
	if num%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

func positiveCheck(num int) {
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

}
