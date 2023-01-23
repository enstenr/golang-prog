package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)
func main (){
	godotenv.Load()
	apiKey:=os.Getenv("API_KEY")
	if apiKey==""{
		log.Fatalln("Missing API_KEY in environment variable")

	}
	ctx:=context.Background()
	client:=gpt3.NewClient(apiKey)

	response,err:=client.Completion(ctx,gpt3.CompletionRequest{
		Prompt:[]string{"The first thing you should need to know about go lang"},
		MaxTokens:gpt3.IntPtr(30),
		Stop:[]string{","},
		Echo: true,

	})
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(response.Choices[0].Text)
}