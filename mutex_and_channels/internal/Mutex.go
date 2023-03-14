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
	runtime.GOMAXPROCS(1)

	wgMutex.Add(2)

	go increment("A")
	go increment("B")

	fmt.Println("Wainting finish.")
	wgMutex.Wait()

}

func increment(id string) {
	defer wgMutex.Done()

	for count := 0; count < 4; count++ {
		mutexMutex.Lock()
		value := couterMutex

		value++

		fmt.Println("ID : ", id, "=> ", value)
		couterMutex = value

		mutexMutex.Unlock()
	}
}
