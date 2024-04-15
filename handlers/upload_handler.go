package handlers

import (
	"Workspace/go-projects/go-chatbot/common"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (ai *GeminiChatHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading image...")

	text := r.FormValue("text")
	// Parse the image from the request
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to parse image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the image to a file
	fileName := "images/" + common.GenerateUniqueFileName() + ".jpg"
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to create image file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to save image to file", http.StatusInternalServerError)
		return
	}

	log.Println("Image saved successfully")

	ctx := context.Background()
	imageBytes, err := common.FetchImageBytes(fileName)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to fetch image bytes", http.StatusInternalServerError)
		return
	}

	responses, err := ai.GeminiAIChatbot.GuessImage(ctx, imageBytes, text)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Send the response back to the client
	responseData := map[string][]string{"responses": responses}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
