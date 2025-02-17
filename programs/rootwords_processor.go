package programs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessRootWords() {
	fmt.Println("\nRoot Word Replacement Program")
	fmt.Println("-----------------------------")

	reader := bufio.NewReader(os.Stdin)

	// Get root words
	var rootWords []string
	for len(rootWords) == 0 {
		fmt.Println("Enter root words (space-separated):")
		rootWordsInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Clean and split the input
		rootWordsInput = strings.TrimSpace(rootWordsInput)
		if rootWordsInput == "" {
			fmt.Println("Error: Root words cannot be empty. Please try again.")
			continue
		}

		rootWords = strings.Fields(rootWordsInput)
		if len(rootWords) == 0 {
			fmt.Println("Error: No valid root words entered. Please try again.")
			continue
		}
	}

	// Get sentence
	var sentence string
	for sentence == "" {
		fmt.Println("\nEnter the sentence:")
		sentenceInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Clean the input
		sentence = strings.TrimSpace(sentenceInput)
		if sentence == "" {
			fmt.Println("Error: Sentence cannot be empty. Please try again.")
			continue
		}
	}

	// Process the sentence
	result := replaceWithRootWords(rootWords, sentence)

	// Display results
	fmt.Println("\nResults:")
	fmt.Println("--------")
	fmt.Printf("Root words: %v\n", rootWords)
	fmt.Printf("Original sentence: %s\n", sentence)
	fmt.Printf("Modified sentence: %s\n", result)
}

func replaceWithRootWords(rootWords []string, sentence string) string {
	words := strings.Fields(sentence)
	result := make([]string, len(words))

	for i, word := range words {
		replaced := false
		wordLower := strings.ToLower(word)

		// Check each root word
		for _, root := range rootWords {
			rootLower := strings.ToLower(root)

			// Check if the word starts with the root
			if strings.HasPrefix(wordLower, rootLower) {
				result[i] = root
				replaced = true
				break
			}
		}

		// If no replacement found, keep original word
		if !replaced {
			result[i] = word
		}
	}

	return strings.Join(result, " ")
}
