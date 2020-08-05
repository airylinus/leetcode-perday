package main

import "fmt"

func main() {
	raw := []int{1, 6, -8, 22, 3, 2, 9, 44, 22}
	PopSort(raw)
	fmt.Println(raw)
}

// PopSort returns sorted
func PopSort(in []int) {
	exchanged := true
	// for exchanged {
	// i := 0
	for i := 0; (i < len(in)-1) && exchanged; i++ {
		exchanged = false
		for j := 0; j < len(in)-1-i; j++ {
			if in[j] > in[j+1] {
				tmp := in[j+1]
				in[j+1] = in[j]
				in[j] = tmp
				exchanged = true
			}
		}
	}
	// }
	fmt.Println(in)
}
