package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"fs1n/llmhelper/bridge"
	"fs1n/llmhelper/config"
	"fs1n/llmhelper/method"
)

func init() {
	err := config.InitConfig()
	if err != nil {
		panic("Init config failed")
	}
	bridge.InitLLMClient(config.Mode_GLM4Flash)
}

func main() {
	story, err := method.GetStoryByWords(context.Background(), []string{"apple", "water", "wink", "black", "drink"})
	if err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.Infoln(fmt.Sprintf("story is [%v]", story))
}
