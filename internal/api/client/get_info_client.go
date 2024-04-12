package client

import (
	"Eff_Mob/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type GetInfoClient struct {
	ReGNum string `json:"regNum"`
}

func Client(log *slog.Logger, regNum []string) []models.CarInfo {
	var infoResponse []models.CarInfo
	for i := range regNum {
		client := http.Client{}

		// Get url for client from .env
		//urlBody, exists := os.LookupEnv("CLIENT_URL")
		//if !exists {
		//	log.Debug("set CLIENT_URL env variable")
		//}

		//url := urlBody + regNum[i]

		body := []byte{}

		url := "http://localhost:8082/info?regNum=" + regNum[i]
		fmt.Println(regNum[i])
		fmt.Println(url)

		//Setup request
		req, err := http.NewRequest(
			"GET", url, bytes.NewBuffer(body))
		if err != nil {
			log.Debug("failed to send request", err)

		}
		// Get response
		resp, err := client.Do(req)
		if err != nil {
			log.Debug("failed to get response", err)

		}
		defer resp.Body.Close()

		// Stdout info:
		fmt.Println("response Status: ", resp.Status)
		fmt.Println("response Body: ", resp.Body)

		// Work with response
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
		}
		respStruct := models.CarInfo{}
		err = json.Unmarshal(responseBody, &respStruct)
		if err != nil {
			fmt.Println("Error unmarshalling response:", err)
		}
		infoResponse = append(infoResponse, respStruct)
	}
	return infoResponse
}
