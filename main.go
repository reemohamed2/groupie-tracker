package main

import (
	"fmt"
	pkg "groupie/pkg"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pkg.MainPage)
	mux.HandleFunc("/artistInfo", pkg.ArtistPage)
	mux.HandleFunc("/feedback", pkg.FeedbackPage)
	// http.HandleFunc("/locationInfo", pkg.LocationPage)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
