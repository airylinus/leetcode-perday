package main

import "fmt"

func main() {
	// i := []int{1, 2, 3, 9, -10, 3, 44, 98, 5}
	// var data map[int]int
	// data["a"] := "a123"
	// QSort(i, 0, len(i)-1)
	// fmt.Println(i)
	//a, b := 1, 333
	//fmt.Println(a)
	a := []int{15, 2, 13}
	QSort(a, 0, len(a)-1)
	fmt.Println(a)
}

// QSort ...
func QSort(in []int, l, r int) {
	if r-l < 2 {
		return
	}
	m := l + (r-l)/2
	//in[0], in[middle] = in[middle], in[0]
	//l++
	for l < r {
		if (in[l] > in[m]) && (in[r] < in[m]) {
			in[l], in[r] = in[r], in[l]
			l++
			r--
			continue
		}
		if in[l] < in[m] {
			l++
		}
		if in[r] > in[m] {
			r--
		}
	}
	//in[0], in[l] = in[l], in[0]
	QSort(in, 0, m)
	QSort(in, m, len(in)-1)
	return
}
