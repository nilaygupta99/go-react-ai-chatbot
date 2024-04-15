package models

import (
	"Workspace/go-projects/go-chatbot/common"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
	Client *openai.Client
}

func NewOpenAIModel(apiKey string) (*OpenAIModel, error) {
	fmt.Println("Creating new OpenAI Model")
	client := openai.NewClient(apiKey)
	return &OpenAIModel{Client: client}, nil
}

func (o *OpenAIModel) GenerateImageFromText(text string) (string, error) {
	fmt.Println("Generating image from text...", text)
	ctx := context.Background()

	// Example image as base64
	reqBase64 := openai.ImageRequest{
		Prompt:         text,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := o.Client.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return "", err
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return "", err
	}

	fileName := "images/" + common.GenerateUniqueFileName() + ".png"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return "", err
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return "", err
	}

	fmt.Println("The image was saved as ", fileName)
	return fileName, nil
}
