package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page %d", sum))
}

func addValue(x, y int) int {
	return x + y
}

func Devide(w http.ResponseWriter, r *http.Request) {
	f, err := devideValue(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot devide by 0")

		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f devided by %f is %f", 100.0, 10.0, f))
}

func devideValue(x, y float32) (float32, error) {

	if y == 0 {
		err := errors.New("cannot devide by 0")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Devide)

	_ = http.ListenAndServe(portNumber, nil)

}
