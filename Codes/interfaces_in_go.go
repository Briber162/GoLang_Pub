/*
 INTERFACES
*/
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

// Used receiver
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	MyCircle := circle{
		radius: 4,
	}
	ansArea := MyCircle.area()
	ansP := MyCircle.perimeter()
	fmt.Println(ansArea, ansP)

	MyRec := rect{
		width:  3,
		height: 5,
	}
	nsArea := MyRec.area()
	nsP := MyRec.perimeter()
	fmt.Println(nsArea, nsP)

}



/* Example 2*/

package main

import "fmt"

type shape interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func callFunc(s shape) {
	ans := s.area()
	ans1 := s.perim()
	fmt.Println(ans, ans1)
}

func main() {

	myCircle := circle{
		radius: 5.0,
	}
	callFunc(myCircle)

}







