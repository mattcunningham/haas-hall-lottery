package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	RED    = "\u001B[31m"
	YELLOW = "\u001B[33m"
	GREEN  = "\u001B[32m"
	END    = "\u001B[0m"
)

// Returns a string formatted to be a specific color when printed in the console
// The only available colors are listed as constants. Add more as necessary.
// ex. formatPrint("my string", RED)
func formatPrint(s string, color string) string {
	return color + s + END
}

// Writes a 404 message to the page
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Println(formatPrint("REQUEST[404]: ", RED), r.URL)
	w.Write([]byte("404 not found"))
	// a custom 404 page can be written from here
}

// main function to serve the site
func main() {
	// serves for the webpage "/" — i.e. the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // we want "/" with no appending text
			NotFound(w, r)
			return
		}
		s := "welcome to the page; we don't know what we're doing"
		w.Write([]byte(s))                                       // prints to webpage
		fmt.Println(formatPrint("REQUEST[200]: ", GREEN), r.URL) // prints to console
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
