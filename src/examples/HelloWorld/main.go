package main

import "fmt" // The input/output (I/O) package

func main() {
	log("Hello World!")

	var summary int = sum(5, 6)
	fmt.Printf("5+6 Sum: %v\n", summary)
	fmt.Printf("5 Factorial: %v\n", factorial(5))
}

// I created a log function just like the NodeJS's console.log
func log(message string) {
	fmt.Printf("%v\n", message)
}

func sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
