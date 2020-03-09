package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Amiibo []struct {
		AmiiboSeries string `json:"amiiboSeries"`
		Character    string `json:"character"`
		GameSeries   string `json:"gameSeries"`
		Head         string `json:"head"`
		Image        string `json:"image"`
		Name         string `json:"name"`
		Tail         string `json:"tail"`
		Type         string `json:"type"`
	} `json:"amiibo"`
}

func getJson() {
	// get http example
	response, err := http.Get("https://www.amiiboapi.com/api/amiibo/?name=mario")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var info Info

	// decode json to struct
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

}

func main() {

	getJson()

}
