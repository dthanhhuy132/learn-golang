package main

import "fmt"

// func main() {
// 	fmt.Println("Hello word con bo")

// 	var whatToSay string
// 	var i int

// 	whatToSay = "Good bye, cruel world"
// 	i = 7

// 	fmt.Println("whatToSay:____", whatToSay)
// 	fmt.Println("i:____", i)

// 	// var sayThing string = saySomeThing() //
// 	sayThing, _ := saySomeThing()

// 	fmt.Println("sayThing:____", sayThing)

// }

// func saySomeThing() (string, string) {
// 	return "something", "else"
// }

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
}

func main() {
	// var myString string
	// myString = "Green"

	// fmt.Println("myString before: ___", myString)
	// changeToRed(&myString)

	user := User{
		FirstName: "Huy",
		// LastName:    "Dona",
		PhoneNumber: "243434",
	}

	fmt.Println("user after: ___", user)

}

// func changeToRed(s *string) {
// 	var newValue string = "red"
// 	*s = newValue
// }
