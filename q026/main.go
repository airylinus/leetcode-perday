package main

import (
	"fmt"
)

func removeDuplicates(nums []int) ([]int, int) {
	// nums := [1, 2, 3, 3, 4]
	p := len(nums) - 1
	fmt.Printf("%p p", nums)
	for p > 0 {
		fmt.Println("p : ", p)
		if nums[p] == nums[p-1] {
			nums = removeAt(nums, p)
		}
		p = p - 1
	}
	fmt.Printf("%p answer: %v \n", nums, nums)
	return nums, len(nums)
}

func removeAt(input []int, index int) []int {
	target := input[:0]
	for k := range input {
		if k != index {
			target = append(target, input[k])
		}
	}
	return target
}

func main() {
	// nums := make([]int, 0)
	nums := []int{1, 1, 3, 3, 4, 4}
	fmt.Printf("%p input: %v \n", nums, nums)
	// fmt.Println(reflect.SliceOf())
	nums, _ = removeDuplicates(nums)
	// fmt.Println(n)

	fmt.Printf("output 2 : %v \n", nums)
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	l, r := 0, 1
	for r < n {
		if nums[l] < nums[r] {
			l++
			nums[l] = nums[r]
		}
		r++
	}
	nums = nums[:l+1]
	return l + 1
}
