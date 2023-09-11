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
	locAndDate := data.GetLocationsAndDates()

	tmp, err := template.ParseFiles("./ui/html/home.html")
	logError(w, err, http.StatusInternalServerError)

	if r.Method == http.MethodGet {
		// get full page with artists
		ans := map[string]interface{}{
			"Artists": search,
			"Search":  search,
			"Filters": locAndDate.Items,
		}
		err = tmp.Execute(w, ans)
		logError(w, err, http.StatusInternalServerError)

	} else if r.Method == http.MethodPost {
		// searching
		datas := r.FormValue("searchInput")

		result := pkg.Search(datas)
		locations := data.GetLocations().Index

		ans := map[string]interface{}{
			"Search":    search,
			"Artists":   result,
			"Locations": locations,
			"Filters":   locAndDate.Items,
		}
		err = tmp.Execute(w, ans)
		logError(w, err, http.StatusInternalServerError)
	} else {
		logError(w, err, http.StatusMethodNotAllowed)
	}
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	artisId := r.URL.Path[len("/artist/"):]
	if artisId == "" || artisId[0] == '0' {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(artisId)
	logError(w, err, http.StatusBadRequest)
	artist, locADate, err := data.GetData(id)
	logError(w, err, http.StatusBadRequest)
	location, err := data.GetLocationById(id)
	logError(w, err, http.StatusBadRequest)
	coor := data.GetCoordinatesBatch(location)
	logError(w, err, http.StatusBadRequest)

	tmp, err := template.ParseFiles("./ui/html/artist-page.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	// fmt.Println(coor)
	ans := map[string]interface{}{
		"Artist":   artist,
		"LocADate": locADate,
		"Location": coor,
	}

	err = tmp.Execute(w, ans)
	logError(w, err, http.StatusInternalServerError)
}
