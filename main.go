// package == project, can have many related files
// all files in a package have to declare what package they belong to
// this file belongs to a package called "main"
// when we run `go build main.go` it creates the "main" package
// because we used "main" here
// The name "main" means that this is an executable package
// Any other name would be a reusable package
package main

// Double quotes are important!
// import gives our package access to some code written in
// another package (in this case, "fmt")
// "fmt" is a package included with Go by default, short for "format"
// printing to the terminal
import "fmt"

// Any package named "main" must have a "main" function in it
// func == function
// func is keyword, then name of function
func main() {
	fmt.Println("Hi there!")
}
