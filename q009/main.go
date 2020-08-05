package main

import (
	"fmt"
)

func main() {
	a := 123454321
	b := 111222
	c := 11
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
	fmt.Println(isPalindrome(c))
}

func isPalindrome(x int) bool {
	nums := make([]int, 0)
	odds := x
	if odds < 0 {
		return false
	}
	if odds < 10 {
		return true
	}
	for odds > 0 {
		nums = append(nums, odds%10)
		odds /= 10
	}
	fmt.Print(nums)
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}
	return true
}
