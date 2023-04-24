package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, erro := fmt.Fprintf(w, "Hello, world!")

		if erro != nil {
			fmt.Println("err", erro)
		}
		fmt.Println(fmt.Sprintf("n bytes written: %d", n))
	})

	http.ListenAndServe(":8080", nil)
}
