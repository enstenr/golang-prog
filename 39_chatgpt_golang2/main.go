package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

func GetResponse(client gpt3.Client,ctx context.Context,question string){
	err:=client.CompletionStreamWithEngine(ctx,gpt3.TextDavinci003Engine,gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		 
			MaxTokens: gpt3.IntPtr(3000),
			Temperature:gpt3.Float32Ptr(0),
		}, func (response *gpt3.CompletionResponse){
			fmt.Println(response)
			fmt.Println(response.Choices[0].Text)
		})
		if err != nil{
			fmt.Print(err)
			os.Exit(13)
		}
		fmt.Printf("\n")
	}
		
type NullWriter int 
func (NullWriter) Write([] byte) (int, error){return 0,nil}
func main() {
	log.SetOutput(new(NullWriter))
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API_KEY in environment variable")

	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "chat with Chat GPT in console",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Println("Say something to quit ('quit')")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true

				default:
					GetResponse(client, ctx, question)
				}
			}
		},
	}
	rootCmd.Execute()
}
