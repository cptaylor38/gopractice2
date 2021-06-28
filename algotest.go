package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// callRobots()
	// rangeExample("Hello")
	// // fizzBuzz(33)
	// fmt.Println("Enter a word: ")
	// var userWord string
	// fmt.Scanln(&userWord)
	// palindrome(userWord)
	// httpExample()
	// pointerExample()
	// variadicTest()
	// fmt.Println(multiply(323, 23092))
	// multipleReturns()
	// interpolation()
	// Shape()
	// to run Shape() func from other file, go must run with both files
	// go run algotest.go interfaceExample.go
	// GoRoutineMain()
	// ChannelsMain()
	httpMain()
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
	// https://golang.org/pkg/io/ioutil/#ReadAll
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

func pointerExample() {
	test := "testing pointer"
	pointerHelper(&test)
}

func pointerHelper(test *string) {
	//* dereference needs to be appended to argument type for compiler
	fmt.Println(*test)
	fmt.Println("location of pointer without dereference: ", test)
	*test = "Pointer tested"
	fmt.Println("Updated: ", *test)

}

func variadicTest() {
	variadicHelper("Some text here", 1, 2, 3, 4, 5)
}

func variadicHelper(msg string, values ...int) {
	fmt.Println(values, msg)
}

func multiply(firstNum, secondNum int) int {
	return firstNum * secondNum
}

func multipleReturns() {
	firstMsg, err := returnHelper("drm")
	secondMsg, err := returnHelper("rrr")
	// drm, rrr prints "wrong acronym".
	// rrr, drm prints " ", "doe ray me"
	// drm, drm prints "doe ray me", "doe ray me"
	// rrr, rrr prints "wrong acronym"
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(firstMsg)
	fmt.Println(secondMsg)
}

func returnHelper(msg string) (string, error) {
	if msg != "drm" {
		return "", fmt.Errorf("Wrong acronym")
	}
	return "doe ray me", nil
}

func interpolation() {
	name := "bob"
	age := 21
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message)
}

func greeterExample() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

	//Declaring greeter struct,
	// then calling the greeter function with g. as a method call.

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	//g greeter makes it a method, because it's referencing the greeter struct?
	//a method is just a function executing in a known context.
	//a known context in go is any type.
	//structs are most common but any types can be used.
	//when greet method is called, greet method gets copy of greeter object that
	// gets named g in context of this method.
	//allows us to access fields on greeter object?????????????????

	//greater object gets copied for method/function to use, I guess.

	//(g *greeter) greet() Pointer/Receiver.

	fmt.Println(g.greeting, g.name)
}

func secondMain() {
	var w Writer = ConsoleWriter{}
	//w variable is holding a Writer which is something that implements writer
	// interface.
	// Writer is the interface, and it could be applied to ConsoleWriter instance,
	// or an instance of something else that performs the same writing action
	// thus polymoprhic behavior
	w.Write([]byte("Hello Go!"))

	// Implicit implementation - if you need to wrap a concrete type and
	// someone hasn't published an interface, you can create an interface
	// that their type implements.

	// https://golang.org/pkg/database/sql/
	// Information on sql related database info for Go? Structs, "Concrete types", methods?
	// Sending statements/making queries all go through the concrete DB object - SQL
	// How to test statements/making queries going through the DB object without a database?
	// Test without a database is to create an interface that replicates the method signature
	// and the DB object from the SQL package will automatically implement it.
	// Don't have to worry about creating interfaces at design time, because consumers
	// of library or api can always create interfaces later? And their interfaces
	// can be shaped to exactly what they need for their application.

}

//Interfaces
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

type ConsoleWriter struct{}

//Method on ConsoleWriter with same name as Write from Writer interface.
//Accepting slice of bites, returning integer and error.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
