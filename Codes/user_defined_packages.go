//mylib.go file

// Package mylib provides some utility functions.
package mylib

import "fmt"

// SayHello prints a greeting message.
func SayHello() {
    fmt.Println("Hello from mylib package!")
}



//main.go file

package main

import "mylib" // Import your user-defined package

func main() {
    mylib.SayHello() // Use a function from the imported package
}
