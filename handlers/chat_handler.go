package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"Workspace/go-projects/go-chatbot/models"
)

type GeminiChatHandler struct {
	GeminiAIChatbot *models.GeminiAIChatbot
	ImageStore      *models.ImageStoreService
}

func NewGeminiChatHandler(geminiAIChatbot *models.GeminiAIChatbot) *GeminiChatHandler {
	return &GeminiChatHandler{
		GeminiAIChatbot: geminiAIChatbot,
	}
}

// Handler function for chat requests
func (ai *GeminiChatHandler) ChatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ChatHandler called....")

	// Parse the request body
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		fmt.Println("Error in decoding request data ", err.Error())
		http.Error(w, "Invalid request bodyyy", http.StatusBadRequest)
		return
	}

	// Get the user's message from the request
	userMessage := requestData["message"]
	fmt.Println("User message received is ", userMessage)

	ctx := context.Background()

	responses, err := ai.GeminiAIChatbot.GenerateTextFromText(ctx, userMessage)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Send the response back to the client
	responseData := map[string][]string{"responses": responses}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
