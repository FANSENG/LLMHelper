package model

import "github.com/openai/openai-go"

type ChatCompletionRequest struct {
	Model             string                         `json:"model"`
	Messages          []openai.ChatCompletionMessage `json:"messages"`
	Stream            *bool                          `json:"stream,omitempty"`
	Temperature       *float64                       `json:"temperature,omitempty"`
	TopP              *float64                       `json:"top_p,omitempty"`
	MaxTokens         *int                           `json:"max_tokens,omitempty"`
	PresencePenalty   *float64                       `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64                       `json:"frequency_penalty,omitempty"`
	Stop              *[]string                      `json:"stop,omitempty"`
	Tools             *[]Tool                        `json:"tools,omitempty"`
	ToolChoice        *string                        `json:"tool_choice,omitempty"`
	RequestId         *string                        `json:"request_id,omitempty"`
	DoSample          *bool                          `json:"do_sample,omitempty"`
	ReturnJSON        *bool                          `json:"return_json,omitempty"`
	Functions         *[]Function                    `json:"functions,omitempty"`
	FunctionCall      *string                        `json:"function_call,omitempty"`
	SystemFingerprint *string                        `json:"system_fingerprint,omitempty"`
	UserInfo          *map[string]string             `json:"user_info,omitempty"`
}

type Tool struct {
	Type     string    `json:"type"`
	Function *Function `json:"function,omitempty"`
}

type Function struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

func NewRequest(model string, messages []openai.ChatCompletionMessage) *ChatCompletionRequest {
	return &ChatCompletionRequest{
		Model:    model,
		Messages: messages,
	}
}
