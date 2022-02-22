package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	major := make(map[int]int)
	for _, i := range nums {
		major[i] += 1
		count++
	}
	for j, k := range major {
		if k > count/2 {
			return j
		}
	}

	return 0
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
