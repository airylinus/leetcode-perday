package m_bytes

import "fmt"

func Simple() {
	var a byte
	a = 97
	fmt.Printf("%c\n", a)
	fmt.Println(a)
}

func CountBit(n uint32) (c uint32) {
	for n != 0 {
		c++
		n = n & (n - 1)
	}
	return c
}
