package main

import (
	// Import http for web framework
	"net/http"
	// Import gin for web framework
	"github.com/gin-gonic/gin"
	// Load environment variables
	"github.com/joho/godotenv"
	// Import fmt for logging
	"fmt"
	// Import os for environment variables
	"os"
	// Import io for reading response body
	"io"
	// Import bytes for creating request body
	"bytes"
	// Import encoding/json for decoding JSON
	"encoding/json"
)

// Define the LLMResponse struct
type LLMResponse struct {
	Response string `json:"response"`
}

// Define the UserRequest struct
type UserRequest struct {
	Message string `json:"message"`
}

// Defining structs for the response body
type Message struct {
    Content string `json:"content"`
}

type Choice struct {
    Message Message `json:"message"`
}

type Response struct {
    Choices []Choice `json:"choices"`
}

func getLLMResponse(c *gin.Context) {
	// Get the request body
	var userRequest UserRequest
	c.BindJSON(&userRequest)

	// Make the request to the Groq API
	url := "https://api.groq.com/openai/v1/chat/completions"
	apiKey := os.Getenv("GROQ_API_KEY")
	model := os.Getenv("GROQ_MODEL")
	data := fmt.Sprintf(`{"model": "%s",
							"messages": [{"role": "user",
							"content": "%s"}]}`, model,userRequest.Message)
	body := bytes.NewBuffer([]byte(data))
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	resp, _ := http.DefaultClient.Do(req)
	

	defer resp.Body.Close()
	responseBody, _ := io.ReadAll(resp.Body)
	var apiResponse Response
    err := json.Unmarshal(responseBody, &apiResponse)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

	// Return the response
	c.IndentedJSON(http.StatusOK,
		LLMResponse{Response: apiResponse.Choices[0].Message.Content})
}


func init() {
	// Groq client initialization
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	router.POST("/LLMResponse", getLLMResponse)
	router.Run("localhost:8080")
}