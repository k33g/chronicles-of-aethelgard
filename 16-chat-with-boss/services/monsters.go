package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Define a struct for your request data
type MonsterRequest struct {
	MonsterName string `json:"monster_name"`
}

func GetMonsterDescription(monsterName string) (string, error) {
	monsterServicesUrl := os.Getenv("MONSTER_SERVICE_HOST")
	monsterDescriptionServiceEndPoint := "/api/monster/generate/description"

	// ---------------------------------------------
	// POST request to the monster description service
	// ---------------------------------------------
	// Create the request data
	requestData := MonsterRequest{
		MonsterName: monsterName,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("😡 Error marshaling JSON: %v\n", err)
		return "", err
	}

	// Create the request
	req, err := http.NewRequest("POST", monsterServicesUrl+monsterDescriptionServiceEndPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("😡 Error creating request: %v\n", err)
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client and make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("😡 Error making request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("😡 Error reading response: %v\n", err)
		return "", err
	}

	// Print the response
	//fmt.Printf("Status Code: %d\n", resp.StatusCode)
	//fmt.Printf("Response Text: %s\n", string(body))

	return string(body), nil
}
