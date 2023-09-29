/*
Strucutres and their declaration along with nested structures

*/

package main

import "fmt"

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

func main() {
	fmt.Println("Hello World")
	MyCar := car{
		Make:   "lambo",
		Model:  "gt",
		Height: 23,
		Width:  45,
		FrontWheel: Wheel{
			Radius:   1,
			Material: "Steel",
		},
		BackWheel: Wheel{
			Radius:   1,
			Material: "Steel",
		},
	}
	fmt.Println(MyCar.Make)

}
