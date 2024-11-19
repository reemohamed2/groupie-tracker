package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path != "/" {
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	data := CollectData()
	tmpl, err := template.ParseFiles("templates/template.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 server error", http.StatusInternalServerError)
	}
}
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	artistName := r.URL.Query().Get("name")
	if artistName == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	data := CollectData()
	artistExists := false
	var selectedArtist Data
	for _, artist := range data {
		if artist.A.Name == artistName { //artistName is the one from the url
			selectedArtist = artist
			artistExists = true
			break
		}
	}
	if !artistExists {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/artistPage.html")
	if err != nil {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "500 server error", http.StatusInternalServerError)
	}
}

func FeedbackPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/feedback" {
		http.ServeFile(w, r, "templates/404error.html")
		return
	}
	tmpl, err := template.ParseFiles("templates/feedback.html")
	if err != nil {
		http.ServeFile(w, r, "templates/404error.html")
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.ServeFile(w, r, "templates/404error.html")
	}
}
