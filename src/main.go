package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NetworkData struct {
	IP string `json:"ip"`
}

func main() {
	// SERVER_URL is a required constant
	SERVER_URL := "http://localhost:3000"

	// Makes a get request to the SERVER_URL
	resp, err := http.Get(SERVER_URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Reads the response body

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// Creates a NetworkData struct
	data := NetworkData{}

	// Unmarshals the response body into the NetworkData struct
	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}
	// Prints the response body
	fmt.Println("Your public IP address is:", data.IP)
}
