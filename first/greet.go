package first

import "fmt"

// Greet this is my first lib package
func Greet() {
	fmt.Println("Hello, good morning!")
	// we can invoke it to method which starts with uppercase and lets run the method
	subFunction()
}

func Message(){
	fmt.Println("its go lang")
	subFunction()
	secondSubFunction()
}

// subFunction is a private method not to use in all files
func subFunction(){
	fmt.Println("Inside sub function")
}

func secondSubFunction(){
	fmt.Println("Second inside sub function")
}