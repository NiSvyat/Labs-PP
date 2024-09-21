package main

import (
	"fmt"
	"math"
)

type Person struct {
	firstName  string
	lastName   string
	middleName string
	age        int
}

func (p Person) getPersonInfo() {
	fmt.Println("Имя: ", p.firstName)
	fmt.Println("Фамилия: ", p.lastName)
	fmt.Println("Отчество: ", p.middleName)
	fmt.Println("Возраст: ", p.age)
}

func (p *Person) birthday() {
	p.age++
}

type Circle struct {
	radius float64
}

// Фигуры реализуют интерфейс, если у них будут метод Area
type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

// Метод Area для Rectangle
func (anyRectanagle Rectangle) Area() float64 {
	return anyRectanagle.height * anyRectanagle.width
}

// Метод Area для Circle
func (anyCircle Circle) Area() float64 {
	return anyCircle.radius * anyCircle.radius * math.Pi
}

func callShapeArea(shapes []Shape) {

	for _, v := range shapes {
		fmt.Println(v.Area())
	}
}

func main() {
	//Номер 1
	firstPerson := Person{
		"Жмышенко",
		"Валерий",
		"Альбертович",
		72,
	}
	firstPerson.getPersonInfo()

	//Номер 2
	firstPerson.birthday()
	firstPerson.getPersonInfo()

	//Номера 3-4
	firstRectangle := Rectangle{height: 2, width: 3}
	fmt.Println(firstRectangle.Area())

	firstCircle := Circle{radius: 3}
	fmt.Println(firstCircle.Area())

	secondRectangle := Rectangle{height: 5, width: 2}

	//Номер 5
	Shapes := []Shape{firstCircle, firstRectangle, secondRectangle}

	fmt.Println("Номер 5")

	callShapeArea(Shapes)

	//6 номер
	var firstBook Stringer = &Book{
		id:    1,
		title: "Звездные войны",
		price: 500}
	firstBook.getInfoBook()
	firstBook.setNewPriceInfoBook(700)
	firstBook.getInfoBook()
}

type Stringer interface {
	getInfoBook()
	setNewPriceInfoBook(int)
}

type Book struct {
	id    int
	title string
	price int
}

func (anyBook Book) getInfoBook() {
	fmt.Println("Id: ", anyBook.id)
	fmt.Println("Название: ", anyBook.title)
	fmt.Println("Цена: ", anyBook.price)
}

func (anyBook *Book) setNewPriceInfoBook(newPrice int) {
	anyBook.price = newPrice
}
