package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	// w.Write([]byte("Hello from Snippetbox"))

	//parse the template using template.ParseFiles() function to read the template file into a template set
	// if an error preent, log the detailed error message and http.Error() function to send a generic 500 ISE error to user

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	//using the template.ParseFiles() to read the files and store the temlates in aset.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// log.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	//use Execute() on the template set to write the templates content as the response body

	err = ts.Execute(w, nil)
	if err != nil {
		// log.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
	}

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		// return
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		// return
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// w.Write([]byte("Create a new snippet..."))

	//creating some variables holding dummy data. we will remove these later on during the build.
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly,slowly!\n\n-Kobayashi"
	expires := "7"

	//pass the data into the SnippetModel.Insert() method, receving the ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//redirect user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snipept?id=%d", id), http.StatusSeeOther)
}
