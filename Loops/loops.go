// -> Loops in GO
// -> works almost the same as any other language out there
// -> Loops in Go doesnt really depend on a condition, you can run an infinite loop and have a return stat to make it stop

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("The Value of i: %v", i)
	}

	// Another example for the plant height...
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("The plant Height right now is: ", plantHeight)
		plantHeight++
	}

	fmt.Println("plant has grown to: ", plantHeight, "inches")

	// Module operator is basically the same and returns a %:
	// % = 12 % 4 = 0
	// % = 16 % 5 = 3

	// AND and OR operators
	true || false // false
	true && false // true
}
