package main

import (
	"fmt"
	"net/http"
	"github.com/dustin/go-humanize"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	fmt.Println(humanize.Bytes(82854982))
	http.ListenAndServe(":3000", nil)
}