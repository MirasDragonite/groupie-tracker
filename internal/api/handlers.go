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
	mux.HandleFunc("/filtered", filter)

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
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		// get full page with artists
		ans := map[string]interface{}{
			"Artists": search,
			"Search":  search,
			"Filters": locAndDate.Items,
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
			"Filters":   locAndDate.Items,
		}
		err = tmp.Execute(w, ans)
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	} else {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
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
	if err != nil {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	artist, locADate, err := data.GetData(id)
	if err != nil {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	location, err := data.GetLocationById(id)

	coor := data.GetCoordinatesBatch(location)
	if err != nil {
		errorHandler(w, http.StatusBadRequest)
		return
	}

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
	if err != nil {
		fmt.Println(err)
		return
	}
}

func filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/filtered" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	locAndDate := data.GetLocationsAndDates()
	search := data.GetArtists()
	locations := data.GetLocations().Index
	minCD, err := strconv.Atoi(r.FormValue("minValueCD"))
	maXCD, err := strconv.Atoi(r.FormValue("maxValueCD"))
	minFA, err := strconv.Atoi(r.FormValue("minValueFA"))
	maxFA, err := strconv.Atoi(r.FormValue("maxValueFA"))
	n1, err := strconv.Atoi(r.FormValue("member1"))
	n2, err := strconv.Atoi(r.FormValue("member2"))
	n3, err := strconv.Atoi(r.FormValue("member3"))
	n4, err := strconv.Atoi(r.FormValue("member4"))
	n5, err := strconv.Atoi(r.FormValue("member5"))
	n6, err := strconv.Atoi(r.FormValue("member6"))
	n7, err := strconv.Atoi(r.FormValue("member7"))
	n8, err := strconv.Atoi(r.FormValue("member8"))
	if err != nil {
		errorHandler(w, http.StatusBadRequest)
		return
	}
	numberOfMembers := make([]int, 0)
	numberOfMembers = pkg.AddToSlice(numberOfMembers, n1, n2, n3, n4, n5, n6, n7, n8)

	locationFromFilter := r.FormValue("locationFromFilter")
	fmt.Println(numberOfMembers)
	filtered := pkg.Filter(minCD, maXCD, minFA, maxFA, locationFromFilter, numberOfMembers)

	tmp, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	ans := map[string]interface{}{
		"Search":    search,
		"Artists":   filtered,
		"Locations": locations,
		"Filters":   locAndDate.Items,
	}

	err = tmp.Execute(w, ans)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
}
