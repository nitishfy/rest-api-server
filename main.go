package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello world")
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("unable to start the server")
	}
}
