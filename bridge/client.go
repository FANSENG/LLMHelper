package bridge

import (
	"fs1n/llmhelper/config"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var LLMClient *openai.Client

func InitLLMClient(model string) {

	if _, ok := config.ModeMap[model]; !ok {
		panic("model not found")
	}

	// new openai client
	LLMClient = openai.NewClient(
		option.WithAPIKey(config.ModeMap[model].APIKEY),
		option.WithBaseURL(config.ModeMap[model].BaseURI),
	)
}
