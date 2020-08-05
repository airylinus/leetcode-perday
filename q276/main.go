package main

import "fmt"

func main() {
	fmt.Print(PaintFence(3, 2))
}

// PaintFence print colors
func PaintFence(n, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k*k
	}
	dp := make(map[int]int, 0)
	dp[2] = k*k
	dp[1] = k
	for i:=2; i < n; i++ {
		dp[i] = dp[i-1]*(k-1)+dp[i-1]*k
	}
	return dp[k]
}