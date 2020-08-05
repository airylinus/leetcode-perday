package main

import "fmt"

// User ...
type User struct {
	Name string
}

func main() {
	var us []User
	fmt.Println(len(us))
	// v := make(chan User, 10)
	// go func() {
	// 	us := []User{
	// 		{Name: "aaa"},
	// 		{Name: "bbb"},
	// 	}
	// 	time.Sleep(3 * time.Second)
	// 	v <- us
	// }()
	// if -31 < -2 {
	// 	fmt.Println("hi")
	// }
	// fmt.Println(<-v)
}
