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
type RoomRequest struct {
	RoomName string `json:"room_name"`
}

func GetRandomRoomName() (string, error) {
	roomServicesUrl := os.Getenv("ROOM_SERVICES_HOST")
	roomNameServiceEndPoint := "/api/room/generate/name"

	// --------------------------------------
	// GET request to the room name service
	// --------------------------------------
	// Make the HTTP GET request
	resp, err := http.Get(roomServicesUrl + roomNameServiceEndPoint)
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

	return string(body), nil
}

func GetRoomDescription(roomName string) (string, error) {
	roomServicesUrl := os.Getenv("ROOM_SERVICES_HOST")
	roomDescriptionServiceEndPoint := "/api/room/generate/description"

	// ---------------------------------------------
	// POST request to the room description service
	// ---------------------------------------------
	// Create the request data
	requestData := RoomRequest{
		RoomName: roomName,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("😡 Error marshaling JSON: %v\n", err)
		return "", err
	}

	// Create the request
	req, err := http.NewRequest("POST", roomServicesUrl+roomDescriptionServiceEndPoint, bytes.NewBuffer(jsonData))
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
