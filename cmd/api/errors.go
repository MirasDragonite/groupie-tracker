package api

import (
	"net/http"
	"os"
)

const (
	pathTo404Page = "./ui/html/404.html"
	pathTo400Page = "./ui/html/400.html"
	pathTo500Page = "./ui/html/500.html"
	pathTO405Page = "./ui/html/405.html"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusNotFound:
		w.WriteHeader(http.StatusNotFound)
		page, err := os.ReadFile(pathTo404Page)
		if err != nil {
			w.Write([]byte("Internal error"))
			return
		}
		w.Write(page)

		return

	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		page, err := os.ReadFile(pathTo400Page)
		if err != nil {
			w.Write([]byte("Internal error"))
			return
		}
		w.Write(page)

		return
	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		page, err := os.ReadFile(pathTo500Page)
		if err != nil {
			w.Write([]byte("Internal error"))
			return
		}
		w.Write(page)

		return
	case http.StatusMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		page, err := os.ReadFile(pathTO405Page)
		if err != nil {
			w.Write([]byte("Internal error"))
			return
		}
		w.Write(page)

		return
	}
}
