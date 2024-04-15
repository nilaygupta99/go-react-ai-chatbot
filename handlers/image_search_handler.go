package handlers

import (
	"Workspace/go-projects/go-chatbot/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIChatHandler struct {
	OpenAIChatbot *models.OpenAIModel
}

func NewOpenAICChatHandler(openAIChatbot *models.OpenAIModel) *OpenAIChatHandler {
	return &OpenAIChatHandler{
		OpenAIChatbot: openAIChatbot,
	}
}

func (o *OpenAIChatHandler) ImageGenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ImageGenHandler called....")

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

	response, err := o.OpenAIChatbot.GenerateImageFromText(userMessage)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Send the base64 encoded image as response
	fileURL := "http://localhost:8080/" + response
	responseData := map[string]string{"image": fileURL}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
