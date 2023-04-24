package main

import "fmt"

type Animal interface {
	Say() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "erman Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Trang",
		Color:         "Red",
		NumberOfTeeth: 190,
	}

	PrintInfo(gorilla)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says :", a.Say(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Say() string {
	return "woft"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d Gorilla) Say() string {
	return "gaooo"
}

func (d Gorilla) NumberOfLegs() int {
	return 213123
}
