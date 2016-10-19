package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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
	s := OpenPage("404.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(s))
}

// Opens a web page in the /static/ directory and returns as a string
// As we aren't dealing with *massive* files, and since this is self-serving
// we don't have to worry about network security (since the program doesn't interact with
// a network).
func OpenPage(fileName string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(wd + "/web/" + fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Main function to serve the site
func main() {
	// serves for the webpage "/" — i.e. the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // we want "/" with no appending text
			NotFound(w, r)
			return
		}
		s := OpenPage("home.html") // the home page content
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(s))                                       // prints to webpage
		fmt.Println(formatPrint("REQUEST[200]: ", GREEN), r.URL) // prints to console
	})

	http.HandleFunc("/lottery", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/lottery" {
			NotFound(w, r)
			return
		}
		s := OpenPage("lottery.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(s))                                       // prints to webpage
		fmt.Println(formatPrint("REQUEST[200]: ", GREEN), r.URL) // prints to console
	})
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:1427/").Start()
	case "windows", "darwin":
		err = exec.Command("open", "http://localhost:1427/").Start()
	default:
		err = fmt.Errorf("unsupported platform!")
	}
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":1427", nil))
}
