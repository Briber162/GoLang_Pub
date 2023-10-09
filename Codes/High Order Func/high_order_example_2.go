// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func add(x, y int) int {
	fmt.Println("ADD")
	return x + y
}

func mul(x, y int) int { return x * y }

func curryFunc(abcd func(int, int) int) func(int, int) int {

	return func(x, y int) int {
		return abcd(x, y)
	}

}

//First Class func : abcd
//High order : curryFunc

func main() {
	funcVar := curryFunc(add)
	funcVar_2 := curryFunc(mul)
	fmt.Println(funcVar(5, 6))
	fmt.Println(funcVar_2(5, 6))

}
