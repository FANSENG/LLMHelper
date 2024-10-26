package model

import (
	"encoding/json"
	"github.com/openai/openai-go"
)

type ChatCompletionResponse struct {
	Content string `json:"content"`
}

type Choice struct {
	Index        int                          `json:"index"`
	Message      openai.ChatCompletionMessage `json:"message"`
	FinishReason string                       `json:"finish_reason"`
	ToolCalls    []ToolCall                   `json:"tool_calls,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ContentFilterResult struct {
	Role  string `json:"role"`
	Level int    `json:"level"`
}

type WebSearchResult struct {
	Icon    string `json:"icon"`
	Title   string `json:"title"`
	Link    string `json:"link"`
	Media   string `json:"media"`
	Content string `json:"content"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

type FunctionCall struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}
