package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie-tracker/internal/data"
)

func Start() {
	host := ":8000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist/", getArtist)

	fmt.Printf("Server loading in port%v\n", host)
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatal("Error executing template:", err)
		return

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	artists := data.GetArtists()

	tmp, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, artists)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	artisId := r.URL.Path[len("/artist/"):]
	if artisId == "" || artisId[0] == '0' {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(artisId)
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	// artist, location, date, err := data.GetArtistById(id)
	artist, locADate, err := data.GetData(id)
	if err != nil {
		fmt.Println("Stupid")
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/artist-page.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
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
