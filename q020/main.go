package main

import "fmt"

func main() {
	fmt.Println(validString("[](){}"))
	fmt.Println(validString("[{()}]"))
	fmt.Println(validString("[]({})"))
	fmt.Println(validString("[{()}]"))
}

func validString(inString string) bool {
	inStr := []rune(inString)
	stack := []rune{}
	i := 0
	for i < len(inStr) {
		stack = append(stack, inStr[i])
		fmt.Println(stack)
		if len(stack) > 1 && isMatch(stack[len(stack)-2], stack[len(stack)-1]) {
			stack = stack[0 : len(stack)-2]
			fmt.Println(stack, " backslace")
		}
		i++
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isMatch(left, right rune) bool {
	if left == 91 && right == 93 {
		fmt.Println(left, right, "match")
		return true
	}
	if left == 123 && right == 125 {
		fmt.Println(left, right, "match")
		return true
	}
	if left == 40 && right == 41 {
		fmt.Println(left, right, "match")
		return true
	}
	fmt.Println(left, right, " not match")
	return false
}
