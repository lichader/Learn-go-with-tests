package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, "Hello, ", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// func main() {
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
// }
