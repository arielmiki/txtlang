package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sashabaranov/go-openai"
)

func main() {
	inputFile := flag.String("i", "", "Input file path")
	flag.StringVar(inputFile, "input", "", "Input file path")
	outputFile := flag.String("o", "", "Output file path")
	flag.StringVar(outputFile, "output", "", "Output file path")
	targetLang := flag.String("l", "", "Target programming language (python or golang)")
	flag.StringVar(targetLang, "lang", "", "Target programming language (python or golang)")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" || *targetLang == "" {
		fmt.Println("Usage: ./txtlang -i <input_file> -l <language> -o <output_file>")
		return
	}

	if *targetLang != "python" && *targetLang != "golang" {
		fmt.Println("Unsupported language. Please use 'python' or 'golang'.")
		return
	}

	txtCode, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	prompt := fmt.Sprintf("Translate the following txtlang code to %s:\n\n%s", *targetLang, string(txtCode))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	translatedCode := resp.Choices[0].Message.Content

	err = ioutil.WriteFile(*outputFile, []byte(translatedCode), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Translation completed successfully. Source: %s, Target language: %s, Output file: %s\n", filepath.Base(*inputFile), *targetLang, *outputFile)
}