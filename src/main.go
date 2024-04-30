package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type NetworkData struct {
	IP string `json:"ip"`
}

type GeoData struct {
	City struct {
		Names struct {
			En string `json:"en"`
		}
	}
	Postal struct {
		Code string `json:"Code"`
	}
	Continent struct {
		Names struct {
			En string `json:"en"`
		}
		Code string
	}
	Subdivisions []struct {
		Names struct {
			En string `json:"en"`
		}
	}
	Country struct {
		Names struct {
			En string `json:"en"`
		}
		IsoCode           string `json:"IsoCode"`
		GeoNameID         int    `json:"GeoNameID"`
		IsInEuropeanUnion bool   `json:"IsInEuropeanUnion"`
	}
	Location struct {
		TimeZone       string  `json:"TimeZone"`
		Latitude       float64 `json:"Latitude"`
		Longitude      float64 `json:"Longitude"`
		MetroCode      int     `json:"MetroCode"`
		AccuracyRadius int     `json:"AccuracyRadius"`
	}
}

type Data struct {
	NetworkData
	GeoData
}

const SERVER_URL = "http://mikelmc.dev:3000"

func main() {

	// Gets argument from command line
	// if --geo or -g is passed, get the geo location on /geo endpoint
	// else, get the public IP on / endpoint

	if len(os.Args) > 1 {
		if os.Args[1] == "--geo" || os.Args[1] == "-g" {
			data := GeoData{}
			getGeoData(&data)
			j, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(j))
			os.Exit(0)
		}
	}
	data := NetworkData{}
	getNetworkData(&data)
	j, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(j))
}

func getGeoData(data *GeoData) {
	resp, err := http.Get(SERVER_URL + "/geo")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}
}

func getNetworkData(data *NetworkData) {
	resp, err := http.Get(SERVER_URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}
}
