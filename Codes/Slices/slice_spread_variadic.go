package main

import "fmt"

func sum(nums ...float64) float64 {
	// ?
	total := 0.0
	for _, ele := range nums {
		total += ele
	}
	return total
}

// don't edit below this line

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

//Function overloading is not allowed in Go thus written diff name
func test1(myS []float64) {
	fmt.Println("Dummy Test funciton")
	fmt.Println(myS)
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)

	//mySlice := {2.0,3.0,4.5}  -- This is not correct slice
	myS := []float64{2.0, 3.0, 4.5} // This is the correct way to declare slice

	testS := []float64{}

	for i := 0; i < 4; i++ {
		testS = append(testS, float64(i))
	}
	fmt.Println(testS)
	
	test(myS...)
	test1(myS)
}
