package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Parameters used to build the API's endpoint.
type Configuration struct {
	TokenAPI string `json:"token_api"`
	CityID   string `json:"city_id"`
	Lang     string `json:"lang"`
	Metrics  string `json:"metrics"`
}

// API's response structure.
type Payload struct {
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp    float64 `json:"temp"`
		TempMin int     `json:"temp_min"`
		TempMax int     `json:"temp_max"`
	} `json:"main"`
	Name string `json:"name"`
}

const endpoint string = "http://api.openweathermap.org/data/2.5/weather"

func main() {

	param, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Configuration
	err = json.Unmarshal(param, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Build the API's endpoint.
	var values = make(url.Values)
	values["id"] = []string{cfg.CityID}
	values["units"] = []string{cfg.Metrics}
	values["appid"] = []string{cfg.TokenAPI}
	values["lang"] = []string{cfg.Lang}

	var url string = fmt.Sprintf("%s?%s", endpoint, values.Encode())

	fmt.Println(url)

	// Get the API's response.
	raw, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Read the API's response.
	res, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Payload
	err = json.Unmarshal(res, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}
