package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	var result int64

	for _, value := range ar {
		result += value
	}

	return result
}

func main() {
	data := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	result := aVeryBigSum(data)
	fmt.Println(result)
}
