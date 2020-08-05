package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 555, 333, 2, 1, 1}
	fmt.Print(removeElement(a, 1))
}


func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}