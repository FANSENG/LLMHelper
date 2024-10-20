package main

import (
	"log"

	"fs1n/llmhelper/bridge"
	"fs1n/llmhelper/config"
	"fs1n/llmhelper/model"
)

func init() {
	err := config.InitConfig()
	if err != nil {
		panic("Init config failed")
	}
	bridge.InitLLMClient(config.Mode_GLM4Flash)
}

func main() {

	resp, err := model.BaseCall(model.NewRequest("glm-4-flash", []model.Message{
		{
			Role:    "user",
			Content: "Please write a 50-word story using apple, water, suck",
		},
	}))
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}
	log.Fatalf("message is: %s", resp.Choices[0].Message.Content)
}
