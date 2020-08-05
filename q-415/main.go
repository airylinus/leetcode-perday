package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "12345567"
	b := "212"
	fmt.Println(addStrings(a, b))
	// ab := byte(a)
	// fmt.Println(ab)
}

func addStrings(num1 string, num2 string) string {
	var i int = len(num1) - 1
	var j int = len(num2) - 1
	var r byte = 0
	var s = []byte{}
	for i >= 0 || j >= 0 || r != 0 {
		if i >= 0 {
			r += (num1[i] - '0')
			i--
		}
		if j >= 0 {
			r += (num2[j] - '0')
			j--
		}

		s = append(s, ((r % 10) + '0'))
		r /= 10
	}
	var n int = len(s) - 1
	for i := 0; i < n; i++ {
		s[i], s[n] = s[n], s[i]
		n--
	}

	return string(s)
}

func add(a, b string) (sum, addon string) {
	na, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
	}
	nb, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
	}
	total := na + nb
	if total < int64(10) {
		sum = fmt.Sprintf("%d", total)
		return
	}
	sum = fmt.Sprintf("%d", total-10)
	addon = "1"
	return
}

// func addStrings(num1 string, num2 string) string {
// 	// len1 := len(num1)
// 	// len2 := len(num2)
// 	// n := len1
// 	// if n < len2 {
// 	// 	n = len2
// 	// }
// 	// for i := 0; i < n; i++ {
// 	// 	s, addon := add(num1[i], num2[i])
// 	// }
// 	return ""
// }
