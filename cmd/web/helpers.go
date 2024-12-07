package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// server error helper writes an error message and stack trace to the errorlog
// then sends a generic 500 Internal Server Error response to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//clienterror elper sends a specific status code and corresponding describes to the user
// use this later for send responses like 400

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//notFound helper , simple a convenienve wrapper around clientError which sends a 404 not found response

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
