// fizzbuzz program

package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 101; i++ { // init i at 1, iterate to 100
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
