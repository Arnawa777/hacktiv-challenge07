package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	ticker := time.NewTicker(15 * time.Second) // set ticker to send request every 15 seconds

	for range ticker.C {
		// generate random data for "water" and "wind" values
		data := map[string]int{
			"water": rand.Intn(100) + 1,
			"wind":  rand.Intn(100) + 1,
		}

		// determine status based on "water" and "wind" values
		waterStatus := ""
		waterValue := data["water"]
		if waterValue < 5 {
			waterStatus = "Safe"
		} else if waterValue >= 5 && waterValue <= 8 {
			waterStatus = "Alert"
		} else {
			waterStatus = "Danger"
		}

		windStatus := ""
		windValue := data["wind"]
		if windValue < 6 {
			windStatus = "Safe"
		} else if windValue >= 6 && windValue <= 15 {
			windStatus = "Alert"
		} else {
			windStatus = "Danger"
		}

		// convert JSON data to bytes
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			continue
		}

		// create HTTP request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		// add header to indicate that the data being sent is JSON
		req.Header.Add("Content-Type", "application/json")

		// create HTTP client and send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		// display response from server
		fmt.Println("Response Status:", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response Body:", string(body))

		// display water and wind status
		fmt.Printf("{\n  \"water\": %d,\n", waterValue)
		fmt.Printf("  \"wind\": %d\n}\n", windValue)
		fmt.Println("Water status:", waterStatus)
		fmt.Println("Wind status:", windStatus)
	}
}
