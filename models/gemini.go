package models

import (
	"Workspace/go-projects/go-chatbot/common"
	"context"
	"fmt"
	"log"

	// "github.com/google/generative-ai-go/genai"
	"cloud.google.com/go/vertexai/genai"
)

type GeminiAIChatbot struct {
	client  *genai.Client
	session *genai.ChatSession
}

func NewGeminiAIChatbot(ctx context.Context, apiKey string, ai_model, project_id, region string) (*GeminiAIChatbot, error) {
	// if apiKey == "" {
	// 	apiKey = os.Getenv("API_KEY")
	// }

	client, err := genai.NewClient(ctx, project_id, region)
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel(ai_model)
	session := model.StartChat()
	if err != nil {
		return nil, err
	}

	fmt.Println("Created a model with ", ai_model)

	return &GeminiAIChatbot{session: session, client: client}, nil
}

func (ai *GeminiAIChatbot) GenerateTextFromText(ctx context.Context, text string) ([]string, error) {
	fmt.Println("Sending message to Gemini bot...")

	resp, err := ai.session.SendMessage(ctx, genai.Text(text))
	if err != nil {
		fmt.Println("Error sending message to Gemini bot..., starting fresh session...", err.Error())
		// try with fresh session
		model := ai.client.GenerativeModel("gemini-pro-vision")
		ai.session = model.StartChat()
		resp, err = ai.session.SendMessage(ctx, genai.Text(text))
		if err != nil {
			return nil, err
		}
	}

	responseTexts := common.GenAICandidatesToString(resp.Candidates)

	return responseTexts, nil
}

func (ai *GeminiAIChatbot) GuessImage(ctx context.Context, jpegData []byte, prompt string) ([]string, error) {
	fmt.Println("Guessing image")

	// reset session since more than one contents isn't supported in gemini-pro-vision
	model := ai.client.GenerativeModel("gemini-pro-vision")
	ai.session = model.StartChat()

	img := genai.ImageData("jpeg", jpegData)
	if prompt == "" {
		prompt = "What does this picture look like? Provide a short answer in less than 8 words."
	}

	fmt.Println("Generating image text with prompt: ", prompt)

	res, err := ai.session.SendMessage(
		ctx,
		img,
		genai.Text(prompt),
	)
	if err != nil {
		return nil, err
	}

	responseTexts := common.GenAICandidatesToString(res.Candidates)

	ai.session = model.StartChat()

	return responseTexts, nil
}
