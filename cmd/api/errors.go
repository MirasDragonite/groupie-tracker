package api

import (
	"net/http"
	"text/template"
)

const (
	pathToErrorPage = "./ui/html/error.html"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	errs := "404"
	switch status {
	case 400:
		errs = "400"
	case 404:
		errs = "404"
	case 405:
		errs = "405"
	case 500:
		errs = "500"
	}

	page, err := template.ParseFiles(pathToErrorPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	err = page.Execute(w, errs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	return
}
