package client

import (
	"Eff_Mob/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func Client(log *slog.Logger, regNum []string) []models.CarInfo {
	var infoResponse []models.CarInfo
	for i := range regNum {
		client := http.Client{}
		urlBody, exists := os.LookupEnv("CLIENT_URL")
		if !exists {
			log.Warn("set CLIENT_URL env variable")
		}
		url := urlBody + regNum[i]
		body := `{"regNum:" + regNum}`
		bodyByte := []byte(body)
		req, err := http.NewRequest(
			"GET", url, bytes.NewBuffer(bodyByte),
		)
		if err != nil {
			log.Warn("failed to send request", err)

		}
		resp, err := client.Do(req)
		if err != nil {
			log.Warn("failed to get response", err)

		}
		defer resp.Body.Close()

		fmt.Println("response Status: ", resp.Status)
		fmt.Println("response Headers: ", resp.Header)
		fmt.Println("response Body: ", resp.Body)
		fmt.Println("response StatusCode: ", resp.StatusCode)

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
