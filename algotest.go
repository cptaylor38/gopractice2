package main

import (
	"fmt"
)

func main() {
	// fizzBuzz(33)
	fmt.Println("Enter a word: ")
	var userWord string
	fmt.Scanln(&userWord)
	palindrome(userWord)
}

func fizzBuzz(num int) {
	for i := 0; i < num; i++ {
		// if i%15 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if i%5 == 0 {
		// 	fmt.Println("Buzz")
		// } else if i%3 == 0 {
		// 	fmt.Println("Fizz")
		// } else {
		// 	fmt.Println(i)
		// }

		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}

func palindrome(word string) {
	s := word[:]
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	if word == reverse {
		fmt.Println(true, "This word is a palindrome.")
	} else {
		fmt.Println(false, "Try again.")
	}
}
