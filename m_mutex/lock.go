package m_mutex

import (
	"fmt"
	"sync"
	"time"
)

var VarTest int32

// FuncTest123 ..
func FuncTest123(VarTest int32) bool {
	VarTest = 4
	if VarTest == 2 {
		return false
	}
	return true
}

func MutexTest() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait()

}
