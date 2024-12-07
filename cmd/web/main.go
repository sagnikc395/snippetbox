package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// DI for storing a holding the application-wide dependencies
// for the web applications.

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
