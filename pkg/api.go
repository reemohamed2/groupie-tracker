package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func ArtistData() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var artistInfo []Artist
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&artistInfo)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return artistInfo
}

func LocationsData() []Locations {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var locationsResp struct {
		Index []struct {
			Locations []string `json:"locations"`
		} `json:"index"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationsResp)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Extract locations from locationsResp
	var locations []Locations
	for _, item := range locationsResp.Index {
		location := Locations{Locations: item.Locations}
		locations = append(locations, location)
	}
	return locations
}

func DatesData() []Dates {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var datesResponse struct {
		Index []struct {
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&datesResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Extract dates from datesResponse
	var datesInfo []Dates
	for _, item := range datesResponse.Index {
		dates := Dates{Dates: item.Dates}
		datesInfo = append(datesInfo, dates)
	}
	return datesInfo
}

func RelationData() []Relation {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var relationResponse struct {
		Index []struct {
			DatesLocations map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&relationResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Extract relations from relationResponse
	var relationInfo []Relation
	for _, item := range relationResponse.Index {
		relation := Relation{DatesLocations: item.DatesLocations}
		relationInfo = append(relationInfo, relation)
	}
	return relationInfo
}

func CollectData() []Data {
	artists := ArtistData()
	locations := LocationsData()
	dates := DatesData()
	relations := RelationData()

	// Ensure all arrays have the same length
	dataData := make([]Data, len(artists))
	for i := 0; i < len(artists); i++ {
		dataData[i].A = artists[i]
		dataData[i].L = locations[i]
		dataData[i].D = dates[i]
		dataData[i].R = relations[i]
	}
	return dataData
}
