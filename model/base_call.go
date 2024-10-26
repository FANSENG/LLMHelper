package model

import (
	"context"

	"github.com/openai/openai-go"

	"fs1n/llmhelper/bridge"
)

func BaseCall(ctx context.Context, request *ChatCompletionRequest) (ChatCompletionResponse, error) {
	var response ChatCompletionResponse

	messages := []openai.ChatCompletionMessageParamUnion{}
	for _, v := range request.Messages {
		if v.Role == "user" {
			messages = append(messages, openai.UserMessage(v.Content))
		} else if v.Role == "assistant" {
			messages = append(messages, openai.AssistantMessage(v.Content))
		}
	}

	charCompletion, err := bridge.LLMClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(request.Model),
	})
	if charCompletion == nil || err != nil {
		return response, err
	}
	response.Content = charCompletion.Choices[0].Message.Content
	return response, nil
}
