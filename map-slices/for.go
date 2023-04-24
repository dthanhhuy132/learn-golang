package main

import "fmt"

func main() {
	// animals := []string{"dog", "chicken", "cat"}
	// for i, amianimal := range animals {
	// 	// for _, amianimal := range animals { => if not using i, it means that dont care about i
	// 	fmt.Println("i:__:", i)
	// 	fmt.Println("amianimal:__:", amianimal)
	// }

	animals := make(map[string]string)

	animals["name"] = "dog"
	animals["leg"] = "4"

	for i, amianimal := range animals {
		// for _, amianimal := range animals { => if not using i, it means that dont care about i
		fmt.Println("i:__:", i)
		fmt.Println("amianimal:__:", amianimal)
	}
}
