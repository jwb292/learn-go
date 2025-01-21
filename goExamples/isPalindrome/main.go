// Palindrome checker, is string same in reverse?

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	// Convert string to lowercase to make it case-insensitive
	input = strings.ToLower(input)

	// Reverse the string
	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}

	// Check if the original string matches the reversed string
	return input == reversed
}

// Main calling isPalindrome
func main() {
	// Examples to run
	testCases := []string{"Radar", "hello", "MADAM", "level", "GoLang", "deified", "racecar", "Mississippi"}

	for _, test := range testCases {
		if isPalindrome(test) {
			fmt.Printf("'%s' is a palindrome.\n", test)
		} else {
			fmt.Printf("'%s' is not a palindrome.\n", test)
		}

	}
}
