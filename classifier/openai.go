package classifier

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/ethanmick/lime/email"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

type openaiClassifier struct {
	client *openai.Client
}

func NewOpenAIClassifier() Classifer {
	apikey := viper.GetString("openai_apikey")
	if apikey == "" {
		slog.Error("no openai apikey found. Export as LIME_OPENAI_APIKEY=")
	}
	client := openai.NewClient(apikey)
	return &openaiClassifier{client}
}

func (c *openaiClassifier) Classify(e *email.Email) ([]email.Label, error) {
	var emailTypes string
	for k, v := range email.EmailTypes {
		emailTypes += fmt.Sprintf("%s: %s\n", k, v)
	}

	var contentLabels string
	for k, v := range email.Content {
		contentLabels += fmt.Sprintf("%s: %s\n", k, v)
	}

	body, err := e.Body.Decode()
	if err != nil {
		return nil, fmt.Errorf("error decoding email body, ensure it is valid urlencoded base64: %w", err)
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf(SystemPrompt, emailTypes, contentLabels),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: body,
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error received from chat completion: %w")
	}

	var labels []email.Label
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &labels)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling GPT response of '%s': %w", resp.Choices[0].Message.Content, err)
	}
	return labels, nil
}

const SystemPrompt = `You are an email classifier. Your job is to read the
contents of an email and any aggregated stats for it and classify the type of
email as well as label the contents. Only use the types and labels in the given
sets. Each label has a succint definition for what it applies to. If none
apply, return an empty JSON array. Only pick a single type. Pick as many content
labels that apply.

Email Types:
%s

Content Labels:
%s

Do not explain your label choice. Respond in a JSON array of the labels. Example:
["notification","order_confirmation","reminder"]

ONLY RESPOND WITH A VALID JSON ARRAY. NOTHING ELSE.`
