package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// creating a map to represent the response data
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	// Calls our function we made to format this data as json and send it as HTTP response status 200
	// handle error if any
	if err := writeJSON(w, http.StatusOK, data); err != nil {
		writeJsonError(w, http.StatusInternalServerError, "err.Error()")
	}
}
