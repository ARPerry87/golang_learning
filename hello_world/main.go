package main

import "fmt" 

func main() {
  fmt.Println("Hi there, welcome to your Go project!")
}

// Go run compiles and executes one or more files
// Go build compiles a bunch of go source code files into a binary
// Go fmt formats the source code according to the Go style guidelines
// Go get downnloads the raw source code of a package and installs it
// Go install compiles and installs a package
// Go test runs the tests in a package
// Go vet checks the source code for common mistakes
// Go doc prints the documentation for a package
// Go generate runs a generator command to produce source code

// after running go build main.go you get a binary file called main
// which then allows you to run the program with only ./main

// A package is like a project or a workspace
// 1 app == 1 package
// Packages can contain multiple files, and in each file you have to declare 
// the package name/that the file belongs to at the top of the file
// The main package is the entry point of the program

// there are two types of packages: executable and reusable
// Executable packages generates a file that we can run
// Main in this case is an executable package and do things
// Reusable packages is code used as helpers, and are dependencies or libraries
// Good place to put reusable logic, like functions, structs, interfaces, etc.

// The name of the package is what determines if it's a reusable or executable package
// The name 'main' is reserved for the executable package

// Any other name other than main is a reusable package

// Executable packages always have a main function inside of them, no ifs ands or buts

// Importing

// import "fmt" means give my main package access to the fmt package
// fmt is a package that provides formatted I/O functions and is part of the standard Go library
// used to print out to the terminal and is useful in debugging
// without importing, the code will not have access to the fmt package
// this happens with debug package, encoding, math, etc.
// Can also import multiple packages, even those from other engineers, such as a package named 'calculator'
// The offical docs for Go for the packages are at https://golang.org/pkg/

// Functions
// Func is short for function, and they function just like methods or functions in other languages

// Patterns in the files are almost always the same

// 1. package declaration
// 2. import statement
// 3. function declaration
// 4. function body
// 5. return statement (if applicable)