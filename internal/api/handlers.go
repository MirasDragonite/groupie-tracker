package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/pkg"
)

func Start() {
	host := ":8000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist/", getArtist)

	fmt.Printf("Server loading in http://localhost%v/\n", host)
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatal("Error executing template:", err)
		return

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	search := data.GetArtists()

	tmp, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		// get full page with artists
		ans := map[string]interface{}{
			"Artists": search,
			"Search":  search,
		}
		err = tmp.Execute(w, ans)

		if err != nil {
			fmt.Println(err)
			return
		}
	} else if r.Method == http.MethodPost {
		// searching
		datas := r.FormValue("searchInput")

		result := pkg.Search(datas)
		locations := data.GetLocations().Index

		ans := map[string]interface{}{
			"Search":    search,
			"Artists":   result,
			"Locations": locations,
		}
		err = tmp.Execute(w, ans)
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	}
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusNotFound)
		return
	}

	artisId := r.URL.Path[len("/artist/"):]
	if artisId == "" || artisId[0] == '0' {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(artisId)
	if err != nil {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	// artist, location, date, err := data.GetArtistById(id)
	artist, locADate, err := data.GetData(id)
	if err != nil {
		fmt.Println("Stupid")
		errorHandler(w, http.StatusBadRequest)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/artist-page.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	ans := map[string]interface{}{
		"Artist":   artist,
		"LocADate": locADate,
	}
	err = tmp.Execute(w, ans)

	if err != nil {
		fmt.Println(err)
		return
	}
}
