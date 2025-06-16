package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // embedding a struct
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "some_email@email.com",
			zipCode: 12345,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v\n", jim) // %+v will print the struct with the field
}

//structs are similar to maps or hashes in ruby but not quite the same
// if we assign alex to type person, it'll make a new struct and populate 
// it with the values of the original struct, that is to say, make the strings
// default to an empty value, not nil like in JS or Ruby
// We then also did some string interpolation to print the keys of the struct to go with the values
// %+v will print the struct with the field names

// we can also embed one struct into another, which is similar to inheritance in OOP
// but not quite the same, as Go does not have inheritance in the traditional sense
// we can also use pointers to structs, which is similar to references in other languages
