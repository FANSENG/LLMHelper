package method

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/openai/openai-go"
	"github.com/sirupsen/logrus"

	"fs1n/llmhelper/model"
)

var prompt = "Given the word [%v], write a %v-word short story that incorporates this word. The story should be realistic and related to everyday life, helping me to remember the word through context."

func GetStoryByWords(ctx context.Context, wordList []string) (string, error) {
	if len(wordList) == 0 || len(wordList) > 10 {
		logrus.Errorln("wordList is empty or more than 10")
		return "word count error", errors.New("wordList is empty or more than 10")
	}
	words := strings.Join(wordList, ",")
	resp, err := model.BaseCall(ctx, model.NewRequest("glm-4-flash", []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: fmt.Sprintf(prompt, words, GetStoryWordCountByLen(len(wordList))),
		},
	}))
	if err != nil {
		logrus.Errorln(fmt.Sprintf("GetStoryByWords err: %v", err))
		return "call api failed", errors.New("call api failed")
	}
	return resp.Content, nil
}

func GetStoryWordCountByLen(count int) int {
	switch {
	case count < 3:
		return 15
	case count < 6:
		return 30
	default:
		return 60
	}
}
