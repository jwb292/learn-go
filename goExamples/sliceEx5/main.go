/* You have a slice of words, and you want to group them into three categories:

Short words: 3 characters or fewer.
Medium words: 4–6 characters.
Long words: More than 6 characters.
Write a program that processes the words and categorizes them using a helper function.
*/

package main

import "fmt"

func main() {
	// Create a slice of animals
	words := []string{"i", "a", "it", "is", "in", "ant", "big", "cat", "apple", "boy", "car", "dog", "elephant", "house", "journey", "mountain", "ocean", "rainbow", "sunset", "universe"}

	shortWords := make([]string, 0)
	mediumWords := make([]string, 0)
	longWords := make([]string, 0)

	for _, word := range words {
		category := categorizedWords(word)
		switch category {
		case "short":
			shortWords = append(shortWords, word)
		case "medium":
			mediumWords = append(mediumWords, word)
		case "long":
			longWords = append(longWords, word)

		}

	}

	fmt.Println("Short words: ", shortWords, len(shortWords))
	fmt.Println("Medium words: ", mediumWords, len(mediumWords))
	fmt.Println("Long words: ", longWords, len(longWords))

}

func categorizedWords(words string) string {
	length := len(words)
	if length <= 3 {
		return "short"
	} else if length <= 6 {
		return "medium"
	} else {
		return "long"
	}
}
