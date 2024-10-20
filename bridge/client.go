package bridge

import (
	"time"

	"fs1n/llmhelper/config"
	openai "fs1n/llmhelper/helper"
)

var LLMClient *openai.Client

func InitLLMClient(model string) {

	if _, ok := config.ModeMap[model]; !ok {
		panic("model not found")
	}

	// new openai client
	LLMClient, _ = openai.NewClient(&openai.Options{
		ApiKey:  config.ModeMap[model].APIKEY,
		Timeout: 30 * time.Second,
		Debug:   true,
		BaseUri: config.ModeMap[model].BaseURI,
		Model:   model,
	})
}
