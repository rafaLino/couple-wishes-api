package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
	"google.golang.org/api/option"
)

type GeminiAdapter struct {
	context context.Context
	model   *genai.GenerativeModel
	ports.AIAdapter
}

func NewGeminiAIAdapter() ports.AIAdapter {
	return &GeminiAdapter{}
}

func (a *GeminiAdapter) Connect() error {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	modelName := os.Getenv("MODEL_NAME")

	if apiKey == "" {
		return errors.New("API_KEY is required")
	}

	if modelName == "" {
		modelName = "gemini-1.5-flash-8b"
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("unexpected error occurred! %v", err)
	}

	argcc := &genai.CachedContent{
		Model: modelName,
		SystemInstruction: genai.NewUserContent(genai.Text(`Your task is to generate json response given a url link. 
			You should only return the json. Do not add any other text or explanation.
			You will receive a url and must generate a json with the following fields: title, description, url and price. 
			the title must be short;
			return empty string if price is nnt found;
			the description must be a small description for the product.
			th result must be in pt-br.
			if you don't find an url, return an empty object.(e.g: "{}")`)),
	}

	cachedContent, err := client.CreateCachedContent(ctx, argcc)

	if err != nil {
		return err
	}

	model := client.GenerativeModelFromCachedContent(cachedContent)
	model.ResponseMIMEType = "application/json"
	a.model = model
	a.context = ctx
	return nil
}

func (a *GeminiAdapter) GenerateResponse(text string) (*entities.WishInput, error) {
	response, err := a.model.GenerateContent(a.context, genai.Text(text))

	if err != nil {
		return nil, err
	}

	return formatResponse(response)
}

func formatResponse(response *genai.GenerateContentResponse) (*entities.WishInput, error) {
	var wish *entities.WishInput
	if len(response.Candidates) > 0 {
		for _, part := range response.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				if err := json.Unmarshal([]byte(txt), &wish); err != nil {
					return nil, err
				}
			}
		}
	}
	return wish, nil
}
