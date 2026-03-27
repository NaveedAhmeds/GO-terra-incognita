// Functions
//
// -> specified number of inputs and vice versa..
// -> the variables are passed by value and not by reference.
// -> for instance increament requires you to mandatorily return the value as to avoid the mismatch.
// -> The best part about GO is that it allows you to return multiple return values.
// -> multiple return should be used when muiltiple return values of the same type are beingt returned.
// -> naked return is the best way to deal with these kinds of returns
//

package main

import (
	"fmt"
)

// "+" symbol in string ussually concats....
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// Multiple return values function.
func getNames() (string, string) {
	return "Jhon", "Doe"
}

// some random function.
func validity(age int) (validitytoAdult int, validitytoDrink int, vallidityToDrive int) {
	validitytoAdult = 18 - age
	if validitytoAdult < 0 {
		validitytoAdult = 0
	}
	validitytoDrink = 21 - age
	if validitytoDrink < 0 {
		validitytoDrink = 0
	}
	vallidityToDrive = 20 - age
	if vallidityToDrive < 0 {
		vallidityToDrive = 0
	}
	return
}

// Explicit return function..
func getCords() (x, y int) {
	return 5, 6
}

// Implicit return function..
func getCoords() (x, y int) {
	x = 6
	y = 7

	// Naked return example
	return
}

// Error return example..
func errRet(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.new("Gooo learn math, can't divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println("Lane,", "Functions in gooooo")
	s1 := "Go"
	s2 := "Pher"
	result := concat(s1, s2)
	fmt.Println("The result is: ", result)

	// The undersocre is used to ignore the second return.
	firstName, _ := getNames()
	_, lastName := getNames()
	fmt.Println("Nice to meet you: ", firstName)
	fmt.Println("last name is: ", lastName)

	// Multiples return function test.
	test1, test2, test3 := validity(18)
	fmt.Println("Adult:  ", test1, "Drink: ", test2, "Drive: ", test3)

	// Explicit
	exp1, exp2 := getCords()
	fmt.Println("Explicit: ", exp1, exp2)

	// Implicit
	imp1, imp2 := getCoords()
	fmt.Print("Implicit: ", imp1, imp2)
}
