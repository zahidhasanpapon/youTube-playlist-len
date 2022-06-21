package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
