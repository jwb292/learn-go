/*In this example, we will:

Categorize ages into two groups: "Teens" (ages 13-19) and "Adults" (ages 20 and above).
Append "No Data Available" or -1 for any cases where no age fits the condition.*/

package main

import "fmt"

func main() {
	// Slice of ages
	ages := []int{12, 15, 17, 20, 25, 18, 30, 40, 11, 19}

	// Empty slice for teens and adults
	teen := make([]int, 0)
	adult := make([]int, 0)

	// Create a for loop to iterate through ages
	for _, age := range ages {
		if isTeen(age) {
			teen = append(teen, age)
		} else if isAdult(age) {
			adult = append(adult, age)
		} else {
			teen = append(teen, -1)
			adult = append(adult, -1)
		}
	}
	// Print results
	fmt.Println("Teen ages are: ", teen)
	fmt.Println("Adult ages are: ", adult)

	// adding "No data available" placeholder
	teen = append(teen, -1)
	adult = append(adult, -1)

	// Print with placeholder
	fmt.Println("Teens with placeholder: ", teen)
	fmt.Println("Adults with placeholder: ", adult)
}

func isTeen(age int) bool {
	return age >= 13 && age <= 19
}

func isAdult(age int) bool {
	return age >= 20
}
