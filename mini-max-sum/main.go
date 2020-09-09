package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {

	result := []int32{0, 0}

	for index := range a {
		if a[index] > b[index] {
			result[0] = result[0] + 1
		} else if a[index] < b[index] {
			result[1] = result[1] + 1
		}
	}

	return result
}

func main() {
	data1 := []int32{5, 6, 7}
	data2 := []int32{3, 6, 10}

	result := compareTriplets(data1, data2)

	fmt.Println(result)
}
