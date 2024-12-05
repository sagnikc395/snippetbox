package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//addr as a new command line flag , a default value and some short help text explaining what the flag controls.
	// the value of the flag will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")

	//using the flag.Parse() function to parse the command-line
	// reads in the command line flag value and assigns it to the addr variable. we need to call this "before" we use the addr variable

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//create a file server which will server file  out of the "./ui/static" directory

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//using mux.Handle() to register the file server as the handler all URL paths that start with "/static/"
	// for matching paths, we strip "/static" prefix

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on %s", addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
