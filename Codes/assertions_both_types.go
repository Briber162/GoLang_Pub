package main

import "fmt"

type shape interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height int
	width  int
}

func printParams(s shape) {

	switch itr := s.(type) {
	case rectangle:
		fmt.Println("RECTANLGE", itr.height, itr.width)
	case circle:
		fmt.Println("CIRCLE", itr.radius)
	default:
		fmt.Println("DEFAULT")
	}

	/*rec, ok := s.(rectangle)
	if ok {
		fmt.Println(rec.height, rec.width)
	}

	cir, ok := s.(circle)
	if ok {
		fmt.Println(cir.radius)
	}
  return fmt.Println("Default case in ok type assertion")
  */

}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return float64(r.height * r.width)
}

func (r rectangle) perim() float64 {
	return float64(2 * (r.height + r.width))
}

func callFunc(s shape) {
	ans := s.area()
	ans1 := s.perim()
	fmt.Println(ans, ans1)

	printParams(s)
}

func main() {

	myCircle := circle{
		radius: 5.0,
	}
	callFunc(myCircle)

	/*	myR := rectangle{
		height : 4,
		width : 3,
	}*/
	callFunc(rectangle{
		height: 4,
		width:  3,
	})
	callFunc(circle{
		radius: 4,
	})
}
