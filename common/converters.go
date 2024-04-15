package common

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/vertexai/genai"
)

func GenAICandidatesToString(candidates []*genai.Candidate) []string {
	fmt.Println("candidatesToString")
	var responseTexts []string
	for _, candidates := range candidates {
		for _, part := range candidates.Content.Parts {
			responseTexts = append(responseTexts, fmt.Sprintf("%s", part))
		}
	}
	return responseTexts
}

func EncodeImage(imagePath string) (string, error) {
	fmt.Println("EncodeImage")
	// Read the saved image file
	savedFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer savedFile.Close()

	// Read the image data into memory
	imageData, err := io.ReadAll(savedFile)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	// Encode the image data as Base64
	encodedImage := base64.StdEncoding.EncodeToString(imageData)
	return encodedImage, nil
}

func FetchImageBytes(imagePath string) ([]byte, error) {
	fmt.Println("FetchImageInBytes")
	// Read the saved image file
	savedFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer savedFile.Close()

	// Read the image data into memory
	imageData, err := io.ReadAll(savedFile)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return imageData, nil
}

func GenerateUniqueFileName() string {
	// Generate a timestamp string using the current time
	timestamp := time.Now().Format("20060102_150405") // Format: YYYYMMDD_HHMMSS

	return timestamp
}
