package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	callRobots()
	rangeExample("Hello")
	// fizzBuzz(33)
	fmt.Println("Enter a word: ")
	var userWord string
	fmt.Scanln(&userWord)
	palindrome(userWord)
	httpExample()
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
		fmt.Println("This word is a palindrome.")
	} else {
		fmt.Println("Not a palindrome. Try again.")
	}
}

func rangeExample(word string) {
	for _, v := range word {
		fmt.Println(string(v))
	}
}

func callRobots() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
		// look up Fatal function
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	//Look up ioutil package
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	//"%s" = the uninterpreted bytes of the string or slice
}

func httpExample() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!")) //response writer giving access to the web request
		//and printing out hello go
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error()) //example of error - port could be blocked.
		//we decide to panic because http will just create an error. we create panic
		//for error handling and prevention of the application trying to continue
	}

	//Review and look into templating engines.
}
