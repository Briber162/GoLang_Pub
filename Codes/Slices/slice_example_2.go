package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	//When you define slice as ans := []float{} it does not allcate any memoery , thus we can not use as ans[index] = value , we need to use append function
	max_len := costs[len(costs)-1].day + 1
	ans := make([]float64, max_len)
	for _, ele := range costs {
		ans[ele.day] += ele.value
	}
	return ans
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}
