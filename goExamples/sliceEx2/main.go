/* Create a slice of strings containing the days of the week: "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday".
Create a new slice that contains only the weekdays (Monday to Friday).
Create another new slice that contains only the weekends (Saturday and Sunday).
Append "Holiday" to both the weekdays and weekends slices.
Print all three slices: the original, weekdays, and weekends. */

package main

import "fmt"

func main() {
	// Creat a slice containing the days of the week
	weekSlice := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	weekdaySlice := make([]string, 0)

	weekendSlice := make([]string, 0)

	// Create a slice of only the weekdays
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// For loop to iterate over the weekSlice
	for _, v := range weekSlice {
		if contains(weekdays, v) {
			weekdaySlice = append(weekdaySlice, v)
		} else {
			weekendSlice = append(weekendSlice, v)
		}

	}

	weekdaySlice = append(weekdaySlice, "Holiday")
	weekendSlice = append(weekendSlice, "Holiday")
	fmt.Println("Week days are: ", weekdaySlice)
	fmt.Println("Weekend days are: ", weekendSlice)

}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false

}
