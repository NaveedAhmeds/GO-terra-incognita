// Variables in GOooooooo...

// -> variables are mostly inferred to their types.
// -> less requirment of having to declare the types explicitly
// -> Types:
//
//  -> bool
//  -> string
//  -> Int
//  -> Uint32
//  -> Uint32
//  -> byte
//  -> rune
//  -> float64
//  -> complex128

package main


import "fmt"

// variables declared outside of the function scope must be initialized as var.
var global_var = "Global"

	// 
	//Short Variable Declaration.
	mascot := "Gopher"

	i := 42           // Int
	f := 3.14         // Float
	g := 0.867 * 0.5i // Complex

	// same line variables pretty GOood.
	var1, var2 := 1, "go goes brrr.."
 
	//Variable conversions..
	temperature := 30
	temperatureFloat := float64(temperature)
 
	//Contants (const does not support short name declarations).
	const stubbornGopher = "Doesn't Move"
	// 
	//String formatting not very GGGGooodddd unfortunately [printf, sprintf]
	// -> %v (default formatter)
	// -> %s (Interpolates a String)
	// -> %d (decimal form)
	// -> %f (Interpolates in decimal)
	// 
	//OPERATORS (operations in GOoooopher)
	// -> == equal to
	// -> != not equal to
	// -> < less than
	// -> > greater than
	// -> <= less or equal
	// -> >= great or eual
	// 
    //TESTINGGGGG>>>>>
	fmt.Println(mascot)
	fmt.Print(global_var)
	fmt.Printf("The type of i: %T\n", i)
	fmt.Printf("The type of fL: %T\n", f)
	fmt.Printf("The type of g: %T\n", g)
	fmt.Println(var1, var2)
}
