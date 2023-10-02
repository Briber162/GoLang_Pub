package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	ans := [][]int{} //Need to use append
	for i := 0; i < rows; i++ {
		tempS := []int{}
		for j := 0; j < cols; j++ {
			tempS = append(tempS, i*j)
		}
		ans = append(ans, tempS)
	}
	return ans
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
