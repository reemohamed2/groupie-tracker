package pkg

type Artist struct {
	Id           uint     `json:"id"`           // Backticks added
	Name         string   `json:"name"`         // Backticks added
	Image        string   `json:"image"`        // Backticks added
	Members      []string `json:"members"`      // Backticks added
	CreationDate uint     `json:"creationDate"` // Backticks added
	FirstAlbum   string   `json:"firstAlbum"`   // Backticks added
}

type LocationsResponse struct {
	Index []struct {
		Id        uint     `json:"id"`        // Backticks added
		Locations []string `json:"locations"` // Backticks added
		Dates     string   `json:"dates"`     // Backticks added
	} `json:"index"` // Backticks added
}

type Locations struct {
	Locations []string `json:"locations"` // Backticks added
}

type Dates struct {
	Dates []string `json:"dates"` // Backticks added
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"` // Backticks added
}

type Data struct {
	A Artist
	R Relation
	L Locations
	D Dates
}
