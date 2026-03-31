//-> Error in GO are the most Gooood thing about the language
// -> Support for catching errors while working on the variable assignment is a way to use errors.
// -> erros in go are just string values
// -> Handling errors in code can basically be dealth with putting them under the hood, meaning turning them to string
// -> The underlying type of an error is an interface, as it can always be implemented accordingly.
// -> Type can also be an error as types can satisfies any interface given the face that they have all the implementations for the interface

package main

func main() {
	user, err := getuser()
	if err != nil {
		fmt.println(err)
		return
	}

	// another exampl of it is
	user, err := getuserprofile(user.id)
	if err != nil {
		fmt.println(err)
		return
	}

	//Converting the values to String in the guard class
	func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error){
		cost, err := sendSMS(msgToCustomer)
		if err != Nil{
			return 0.0, err
		}
	}
}
