package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//addr as a new command line flag , a default value and some short help text explaining what the flag controls.
	// the value of the flag will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")

	//using the flag.Parse() function to parse the command-line
	// reads in the command line flag value and assigns it to the addr variable. we need to call this "before" we use the addr variable

	flag.Parse()

	//info log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	//logger for writing error messages in the same way, but use stderr
	// the destination and use the log.Lshortfile flag to include the relevant file name and line number

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//create a file server which will server file  out of the "./ui/static" directory

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//using mux.Handle() to register the file server as the handler all URL paths that start with "/static/"
	// for matching paths, we strip "/static" prefix

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// log.Println("Starting server on %s", addr)
	//print messages using our own logger

	//init a new http.Server struct. we are setting the addr and handler fields that the server uses the same network address
	// and routes as before and the errorlog field so that the server now uses the custom errorlog logger for
	// sendinf the event of any problems.

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	// use new logger
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
