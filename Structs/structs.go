// Structs in GO!!
// -> Key value pairs similarities.
// -> Struct Keys can hold any type
//
//

package main

import "fmt"

type messageToSend struct {
	message     string
	phoneNumber int
}

func test(m messageToSend) {
	fmt.Printf("Sending message as: '%s' to: '%v'", m.message, m.phoneNumber)
}

// key exmaple...
type car struct {
	Make       string
	Model      string
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}


// ------ Anonymous Structs in GO -----
//  -> just like a normal struct but is defined without a name and therefore cannot be referenced elsewhere in the code.
//  -> best use case is to prevent from re-using a struct definition you never intended to use 
//	Note: best coding practice to use general structs rather than the Anonymous struct
myCar := struct{
	Make string
	Model string
}{
	Make: "Mercedes",
	Model: "AMG"
}


//Auth practice to work on the structs with
//-> example
// -> getBasicAuth doesn't exist in this scope although we imahine it to be an authentiacatipon function returning some message after processing.
//
//
	type authenticationInfo struct{
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string{
	fmt.Sprintf("
		"Authorization: Basic %s:%s",
			authI.username,
			authI.passowrd,
		")
}



func main() {
	test(messageToSend{
		message:     "Naveed",
		phoneNumber: 1234567890,
	})

	// Example use case of the holding key mutation.
	myCar := car{}
	myCar.FrontWheel.Radius = 5
}
