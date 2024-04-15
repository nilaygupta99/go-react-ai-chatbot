# go-react-ai-chatbot
AI chatbot using React and GO. Uses Google Vertex AI and OpenAI for various text and image generation

## Setting up frontend server
```
  cd chatbot-frontend
  yarn install
  yarn start
```

## Starting backend server 
```
  go mod tidy && go mod vendor
  go run main.go
```

### Pre-requisites and Setup - 
- Set OpenAPI API key in the environment variable (required for image generation): 
  - ``` export OPENAI_API_KEY=<API_KEY> ```
- Install gCloud CLI (required to interact with Google Vertex AI for text generation from text and image):
  > https://cloud.google.com/sdk/docs/install
- Authenticate using Google Cloud -
  - ``` gcloud auth application-default login ```
- In main.go line 24, set Google Cloud project ID and region
  - https://github.com/nilaygupta99/go-react-ai-chatbot/blob/main/main.go#L24
 
## Resources and links used for reference - 
- https://platform.openai.com/docs/guides/images/image-generation
- https://github.com/sashabaranov/go-openai
- https://ai.google.dev/tutorials/get_started_go#set-up-project

## Screenshots
- https://github.com/nilaygupta99/go-react-ai-chatbot/tree/main/screenshots

<img src="https://github.com/nilaygupta99/go-react-ai-chatbot/blob/main/screenshots/Screenshot%202024-04-15%20204504.png" width="400">
<img src="https://github.com/nilaygupta99/go-react-ai-chatbot/blob/main/screenshots/Screenshot%202024-04-15%20204223.png" width="400">
<img src="https://github.com/nilaygupta99/go-react-ai-chatbot/blob/main/screenshots/Screenshot%202024-04-15%20204826.png" width="400">


