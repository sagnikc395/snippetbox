package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler for showSnippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display a specific snippet ... "))

	//extract the value of the id parameter from the query string and try to convert to a string
	// using strconv.Atoi(). if it can't be converted to an integer, or the value is < 1, we return a 404 not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//using fmt.Fprintf() function to interpolate the id value with our response and write it http.ResponseWriter
	fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)
}

// handler for createSnippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Create a new snippet ... "))
	// send a 405 (method not alowed) HTTP status code unless the request method is POST.
	// and also set the header to Allow: Post header , to let the user know which request methods are supported for that
	// particular URL

	if r.Method != "POST" {
		// w.Header().Set("Allow", "Post")
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// return
		// instead of this much hard work , we can directly send a 405 status code and "Method Allowed" string as the response body
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

// handler function which writes a byte slice containing "hello from snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	//check if current request URL path matches / or not
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

func main() {

	//init a new servermux , then register the home function as the handler for
	// the / URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//register the handlers
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//start a new web server on the http.listenandserve
	log.Println("starting server on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
