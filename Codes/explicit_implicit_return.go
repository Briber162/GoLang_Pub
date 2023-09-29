/* 
Explicit return values : returning each value explicitly in the return statement
Implicit return value : only return statement will return all the values
*/

package main

import (
	"errors"
	"fmt"
)

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div , errors.New("STR")
}

func main() {
	//fmt.Println("Hello, 世界")
	first, second, third := calculator(5, 6)
	fmt.Println(first, second, third)
}
