package internal

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	couterMutex int
	wgMutex     sync.WaitGroup
	mutexMutex  sync.Mutex
)

func Mutex() {
	wgMutex.Add(2)

	go increment("A")
	go increment("B")

	fmt.Println("Wainting finish.")
	wgMutex.Wait()

}

func increment(id string) {
	defer wgMutex.Done()

	for count := 0; count < 2; count++ {
		mutexMutex.Lock()
		value := couterMutex

		runtime.Gosched()
		value++

		fmt.Println("ID : ", id, "=> ", value)
		couterMutex = value

		mutexMutex.Unlock()
	}
}
