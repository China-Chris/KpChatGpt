package services

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
	translates  = gpt3.TextDavinci002Engine
)

var client gpt3.Client

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

func InitClient(api string) {
	client = gpt3.NewClient(api)
}

func GetAnswer(question string, model string, ch chan string) {
	fmt.Print("User: ")
	fmt.Println(question)
	fmt.Print("Bot: ")
	reply := ""
	i := 0
	ctx := context.Background()
	switch model {
	case "QA":
		if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:   gpt3.IntPtr(100),
			Temperature: gpt3.Float32Ptr(0),
			TopP:        gpt3.Float32Ptr(1),
		}, func(resp *gpt3.CompletionResponse) {
			if i > 1 {
				fmt.Print(resp.Choices[0].Text)
				reply += resp.Choices[0].Text
				ch <- resp.Choices[0].Text
			}
			i++
		}); err != nil {
			log.Fatalln(err)
		}
		if reply != "" {
		}
	case "Friend":
		if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:        gpt3.IntPtr(60),
			Temperature:      gpt3.Float32Ptr(0.5),
			TopP:             gpt3.Float32Ptr(1),
			FrequencyPenalty: 0.5,
		}, func(resp *gpt3.CompletionResponse) {
			if i > 1 {
				fmt.Print(resp.Choices[0].Text)
				reply += resp.Choices[0].Text
				ch <- resp.Choices[0].Text
			}
			i++
		}); err != nil {
			log.Fatalln(err)
		}
		if reply != "" {
		}
	case "Chat":
		if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:        gpt3.IntPtr(150),
			Temperature:      gpt3.Float32Ptr(0.9),
			TopP:             gpt3.Float32Ptr(1),
			FrequencyPenalty: 0.6,
		}, func(resp *gpt3.CompletionResponse) {
			if i > 1 {
				fmt.Print(resp.Choices[0].Text)
				reply += resp.Choices[0].Text
				ch <- resp.Choices[0].Text
			}
			i++
		}); err != nil {
			log.Fatalln(err)
		}
		if reply != "" {
		}
	case "FactualAnswering":
		if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:   gpt3.IntPtr(60),
			Temperature: gpt3.Float32Ptr(0.9),
			TopP:        gpt3.Float32Ptr(1),
		}, func(resp *gpt3.CompletionResponse) {
			if i > 1 {
				fmt.Print(resp.Choices[0].Text)
				reply += resp.Choices[0].Text
				ch <- resp.Choices[0].Text
			}
			i++
		}); err != nil {
			log.Fatalln(err)
		}
		if reply != "" {
		}
	case "Translates":
		if err := client.CompletionStreamWithEngine(ctx, translates, gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:   gpt3.IntPtr(60),
			Temperature: gpt3.Float32Ptr(0.9),
			TopP:        gpt3.Float32Ptr(1),
		}, func(resp *gpt3.CompletionResponse) {
			if i > 1 {
				fmt.Print(resp.Choices[0].Text)
				reply += resp.Choices[0].Text
				ch <- resp.Choices[0].Text
			}
			i++
		}); err != nil {
			log.Fatalln(err)
		}
		if reply != "" {
		}
	}
	ch <- "\n"
}
