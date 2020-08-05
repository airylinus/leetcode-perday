package main

import "fmt"

func main() {
	fmt.Println(sumPrefix([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 5))
}

func sumPrefix(arr []int, start, end int) int {
	prefixSum := make(map[int]int)
	for k, v := range arr {
		if k == 0 {
			prefixSum[k] = v
		} else {
			prefixSum[k] = prefixSum[k-1] + v
		}
	}
	return prefixSum[end] - prefixSum[start]
}
