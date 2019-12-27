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
	
	var port = os.Getenv("PORT")
		// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port

	http.ListenAndServe(port, nil)
}