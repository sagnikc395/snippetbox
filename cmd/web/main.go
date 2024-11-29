package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//craete a file server which will server file  out of the "./ui/static" directory

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//using mux.Handle() to register the file server as the handler all URL paths that start with "/static/"
	// for matching paths, we strip "/static" prefix

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
