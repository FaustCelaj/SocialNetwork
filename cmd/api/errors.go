package main

import (
	"log"
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s, error: %s", r.Method, r.URL.Path, err.Error())
	writeJsonError(w, http.StatusInternalServerError, "the server encountered a problem")

}

func (app *application) StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s, error: %s", r.Method, r.URL.Path, err.Error())
	writeJsonError(w, http.StatusBadRequest, err.Error())

}

func (app *application) NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s, error: %s", r.Method, r.URL.Path, err.Error())
	writeJsonError(w, http.StatusNotFound, "not found")

}
