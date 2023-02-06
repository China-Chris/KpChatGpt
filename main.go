package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"log"
)

const (
	maxTokens   = 3000
	temperature = 0.7
	engine      = gpt3.TextDavinci003Engine
	question    = "你好你是谁？"
	apiKey      = "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"
)

func main() {
	fmt.Print("User: ")
	fmt.Println(question)
	client := gpt3.NewClient(apiKey)
	fmt.Print("Bot: ")
	reply := ""
	i := 0
	ctx := context.Background()
	//resp, err := client.Completion(ctx, gpt3.CompletionRequest{
	//	Prompt: []string{"2, 3, 5, 7, 11,"},
	//})
	if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(maxTokens),
		Temperature: gpt3.Float32Ptr(temperature),
	}, func(resp *gpt3.CompletionResponse) {
		if i > 1 {
			fmt.Print(resp.Choices[0].Text)
			reply += resp.Choices[0].Text
		}
		i++
	}); err != nil {
		log.Fatalln(err)
	}
	fmt.Print(reply)
}
