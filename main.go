package main

import (
	"Workspace/go-projects/go-chatbot/handlers"
	"Workspace/go-projects/go-chatbot/middleware"
	"Workspace/go-projects/go-chatbot/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	staticDir := "./images"

	// Create a file server handler for the static directory
	fileServer := http.FileServer(http.Dir(staticDir))

	// Register the file server handler with a route
	http.Handle("/images/", http.StripPrefix("/images/", fileServer))

	ctx := context.Background()
	geminiAI, err := models.NewGeminiAIChatbot(ctx, os.Getenv("API_KEY"), "gemini-pro-vision", "diagflow-test-420214", "us-central1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	geminiChatHandler := handlers.NewGeminiChatHandler(geminiAI)

	openAPIAI, err := models.NewOpenAIModel(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	openAIChatHandler := handlers.NewOpenAICChatHandler(openAPIAI)

	// Define HTTP handler function for handling chat requests
	// Gorilla can also be used for complex HTTP routing
	http.HandleFunc("/chat", geminiChatHandler.ChatHandler)
	http.HandleFunc("/upload", geminiChatHandler.UploadHandler)
	http.HandleFunc("/imagesearch", openAIChatHandler.ImageGenHandler)

	// Start the server
	fmt.Println("Go Chatbot server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.AddCORSHeaders(http.DefaultServeMux)))
}
