package model

import (
	"encoding/json"

	"fs1n/llmhelper/bridge"
	"fs1n/llmhelper/config"
	"fs1n/llmhelper/helper"
)

func BaseCall(request *ChatCompletionRequest) (ChatCompletionResponse, error) {
	var response ChatCompletionResponse
	resp, err := bridge.LLMClient.Post(config.ModeMap[request.Model].URI, helper.StructToMap(request))
	if resp == nil || err != nil {
		return response, err
	}
	err = json.Unmarshal(resp.Raw(), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
