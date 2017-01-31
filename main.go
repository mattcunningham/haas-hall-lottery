package lottery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	//"runtime"
	"strings"
	"os"
	
//	"google.golang.org/appengine"
)

const (
	RED    = "\u001B[31m"
	YELLOW = "\u001B[33m"
	GREEN  = "\u001B[32m"
	END    = "\u001B[0m"
)

type fileError struct {
	error string
}

func (e fileError) Error() string {
	return fmt.Sprintf("%s", e.error)
}

// Returns a string formatted to be a specific color when printed in the console
// The only available colors are listed as constants. Add more as necessary.
// ex. formatPrint("my string", RED)
func formatPrint(s string, color string) string {
	return color + s + END
}

// Writes a 404 message to the page
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatPrint("REQUEST[404]: ", RED), r.URL)
	s, _ := OpenPage("404.html")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(s))
	w.WriteHeader(http.StatusNotFound)
}

// Opens a web page in the /static/ directory and returns as a string
// As we aren't dealing with *massive* files, and since this is self-serving
// we don't have to worry about network security (since the program doesn't interact with
// a network).
func OpenPage(fileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(wd + "/web/" + fileName)
	if err != nil {
		return "", fileError{"File not found."}
	}
	return string(data), nil
}

// Main function to serve the site
func init() {
	// serves for the webpage "/" — i.e. the homepage
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // we want "/" with no appending text
			NotFound(w, r)
			return
		}
		s, _ := OpenPage("home.html") // the home page content
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(s))                                       // prints to webpage
		fmt.Println(formatPrint("REQUEST[200]: ", GREEN), r.URL) // prints to console
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post" && r.Method != "POST" {
			NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		entries := CSVtoEntries(r.FormValue("data"))
		sorted := Sort(entries)
		prioritized := Prioritize(sorted)
		admitted := Admit(prioritized, 3)
		data, err := json.Marshal(admitted) // the error needs to do something
		if err != nil {
			fmt.Println(err)
		}
		w.Write(data)
		fmt.Println(formatPrint("REQUEST[200]: ", GREEN), r.URL) // prints to console
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/static") {
			NotFound(w, r)
		}
		s, err := OpenPage(r.URL.Path[1:])
		if err != nil {
			NotFound(w, r)
			return
		}
		if strings.HasSuffix(r.URL.Path, ".jpg") {
			w.Header().Set("Content-Type", "image/jpg")
		} else if strings.HasSuffix(r.URL.Path, ".html") {
			w.Header().Set("Content-Type", "text/html")
		} else if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		}
		w.Write([]byte(s))
		fmt.Println(formatPrint("REQUEST[200]: ", YELLOW), r.URL)
	})

}
