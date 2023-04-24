package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User) // => change User to interface {} = any as typescript

	// myMap["dog"] = 1
	// myMap["other-dog"] = "Casu"

	me := User{
		FirstName: "Huy",
		LastName:  "Doan",
	}

	myMap["me"] = me

	fmt.Println("myMap:___", myMap["me"].LastName)

	var myNewVar string
	myNewVar = "11.2"
	fmt.Println("myNewVar", myNewVar)

	for i := 1; i <= 10; i++ {
		fmt.Println("i: __", i)
	}
}
