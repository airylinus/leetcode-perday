package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Get the difference set of two sets Example
set1 = []int{1,2,3,333},
set2 = []int{1,3,6,99,-1}
Expect   diff= []int{2, 333, 6, 99 -1}
int64
**/
func main() {
	setA := []int{1, 2, 3, 333, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	setB := []int{1, 3, 6, 99, -1}
	QSort(setA)
	QSort(setB) //
	fmt.Println(setA)
	fmt.Println(setB)
	diff := make([]int, 0)
	for i := 0; i < len(setA); i++ {
		if BSearch(setB, setA[i]) == -1 {
			diff = append(diff, setA[i])
		}
	}
	for i := 0; i < len(setB); i++ {
		if BSearch(setA, setB[i]) == -1 {
			diff = append(diff, setB[i])
		}
	}
	fmt.Println(diff)
}

// func findDiffRaw(a, b []int) []int {

// }

// QSort 快速排序。 N*logN
func QSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := 0
	r := len(nums) - 1
	rand.Seed(time.Now().Unix())
	m := rand.Int() % len(nums)
	nums[m], nums[r] = nums[r], nums[m]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[r] {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
	nums[l], nums[r] = nums[r], nums[l]
	QSort(nums[:l])
	QSort(nums[l+1:])
	return
}

// BSearch 二分查找  logN
func BSearch(haystack []int, niddle int) (idx int) {
	idx = -1
	left := 0
	right := len(haystack) - 1
	for (right - left) > 1 {
		middle := (right + left) / 2
		if haystack[middle] == niddle {
			idx = middle
			return
		}
		if haystack[middle] > niddle {
			right = middle
		} else {
			left = middle
		}
	}
	if right-left == 1 {
		if haystack[right] == niddle {
			idx = right
		}
		if haystack[left] == niddle {
			idx = left
		}
	}
	return
}
