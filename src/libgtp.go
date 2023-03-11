package src

import (
	"fmt"
	"log"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"golang.org/x/net/context"
)

var client gpt3.Client

func InitGTPClient(api string) {
	client = gpt3.NewClient(api)
}

// Sent request to ChatGTP
func GetResponse(question string) string {
	ctx := context.Background()
	resp, err := client.CompletionWithEngine(ctx,
		// "text-davinci-002",
		"text-davinci-003",
		gpt3.CompletionRequest{
			Prompt:    []string{question},
			MaxTokens: gpt3.IntPtr(3000),
		})
	if err != nil {
		log.Fatalln(err)
	}

	// err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
	// 	Prompt: []string{
	// 		quesiton,
	// 	},
	// 	MaxTokens:   gpt3.IntPtr(3000),
	// 	Temperature: gpt3.Float32Ptr(0),
	// }, func(resp *gpt3.CompletionResponse) {
	// 	fmt.Print(resp.Choices[0].Text)
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(13)
	// }
	fmt.Printf("GTP Chat Response\n %v \n", resp.Choices[0].Text)
	return resp.Choices[0].Text
}
