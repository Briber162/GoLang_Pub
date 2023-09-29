/*
  Receivers in GO
*/
package main

import "fmt"

// Define a struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Define a method named "Area" for the Rectangle struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create an instance of the Rectangle struct
    rectangle := Rectangle{Width: 10, Height: 5}

    // Call the method on the struct instance
    area := rectangle.Area()

    // Print the result
    fmt.Println("Area:", area)
}
