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


/* EMBEDED AND NESTED STRUCTS */

package main

import "fmt"

// EMBEDEDD STRUCTS
type car struct {
	model string
	make  string
}

type truck struct {
	tyres int
	car
}

//NESTED STRUCT
type bird struct {
	wings int
	name  string
}

type crow struct {
	myBird bird
	color  string
}

func main() {
	myTruck := truck{
		tyres: 8,
		car: car{
			model: "GT",
			make:  "1987",
		},
	}
	fmt.Println(myTruck.model)

	myC := crow{
		color: "black",
		myBird: bird{
			wings: 2,
			name:  "CROW",
		},
	}

	fmt.Println(myC.color, myC.myBird.name)

}




