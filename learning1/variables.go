package main

// use fmt to handle i/o, like print/echo in other language
import "fmt"

// variable declaration
var sentences = "Hello world"

func calculation(){
	// go is statically typed
	// it cannot use math operator when the data type does not match
	// can declare variable into specific type
	var x int16 = 20
	var y int16 = 30

	fmt.Println(x+y)
}

func main() {
	calculation()
	// println can accept string or other type, Printf used for interpolation
	fmt.Println(sentences)
}


// use go run <filename>.go to execute the file
// use go build to compile the code into binary and excute the .exe file instead