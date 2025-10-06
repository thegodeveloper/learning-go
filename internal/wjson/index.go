package wjson

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func Index(show bool) {
	if show {
		fmt.Println("-- JSON")

		movies := []Movie{
			{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Berdman"}},
			{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
			{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		}

		marshalingStructToJson(movies)

		marshalingStructToJsonFormatted(movies)
	}
}

func marshalingStructToJson(movies []Movie) {
	fmt.Println("-- Marshaling Struct to JSON")
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("Movies: %s\n\n", data)

	unmarshalingDataToStruct(data)
}

func marshalingStructToJsonFormatted(movies []Movie) {
	fmt.Println("-- Marshaling Struct To JSON Formatted")
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("Movies Indented: %s\n\n", data)
}

func unmarshalingDataToStruct(movies []byte) {
	fmt.Println("-- Unmarshaling Data To Struct")
	var titles []struct{ Title string }
	if err := json.Unmarshal(movies, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("Titles: %v\n\n", titles)
}
