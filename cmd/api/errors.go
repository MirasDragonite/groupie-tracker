package api

import (
	"net/http"
	"text/template"
)

const (
	err400          = "400"
	err404          = "404"
	err405          = "405"
	err500          = "500"
	pathToErrorPage = "./ui/html/error.html"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		w.WriteHeader(http.StatusNotFound)
		page, err := template.ParseFiles(pathToErrorPage)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		err = page.Execute(w, err404)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		return
	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		page, err := template.ParseFiles(pathToErrorPage)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		err = page.Execute(w, err400)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		return
	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		page, err := template.ParseFiles(pathToErrorPage)
		if err != nil {
			w.Write([]byte("Internal server error"))
			return
		}
		err = page.Execute(w, err500)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		return
	case http.StatusMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		page, err := template.ParseFiles(pathToErrorPage)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		err = page.Execute(w, err405)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		return
	}
}
