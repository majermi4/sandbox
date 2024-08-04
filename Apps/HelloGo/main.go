package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server listening on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
