package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
}
