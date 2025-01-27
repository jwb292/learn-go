/* Takes a list of ages.
Sorts them into two categories:
Adults: ages 18 and above.
Minors: ages below 18.
Then, we can append "Not Available" to both slices as a final step. */

package main

import "fmt"

func main() {

	// Create list of ages
	ages := []int{15, 22, 17, 19, 14, 25, 12, 30, 16, 18, 19}

	// Create empty slice for minors
	minors := make([]int, 0)

	// Create empty slice for Adults
	adults := make([]int, 0)

	// For loop to iterate over ages
	for _, age := range ages {
		if isAdult(age) {
			adults = append(adults, age)
		} else {
			minors = append(minors, age)

		}
	}

	fmt.Println("Ages of the Adults: ", adults)
	fmt.Println("Ages of the Minors: ", minors)

}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
