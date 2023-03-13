package internal

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func ExemploTwo() {
	runtime.GOMAXPROCS(1)

	at := 5000

	wg1.Add(2)

	go numerosPrimos("A", at)
	go numerosPrimos("B", at)

	fmt.Println("Wait finish.")
	wg1.Wait()
	fmt.Println("Program finish.")
}

func numerosPrimos(id string, at int) {

next:
	for outer := 2; outer < at; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", id, outer)
	}
	wg1.Done()
}
