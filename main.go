package main

import (
	"log"
	"net/http"
)

// handler for showSnippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet ... "))
}

// handler for createSnippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet ... "))
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
