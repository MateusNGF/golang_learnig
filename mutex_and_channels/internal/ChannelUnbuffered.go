package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wgChannelsTwo sync.WaitGroup
)

/*
Channel unbuffered.
*/
func ChannelUnbuffered() {
	baston := make(chan int)
	wgChannelsTwo.Add(1)

	go Runner(baston)

	fmt.Printf("====> Prepare runner ....\n")
	time.Sleep(2 * time.Second)

	baston <- 1

	wgChannelsTwo.Wait()
	fmt.Println("====> Program finish...")
	close(baston)
}

func Runner(baston chan int) {
	var nextRunner int

	runner := <-baston

	fmt.Printf("Runner %d Running with baston\n", runner)

	if runner != 4 {
		nextRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", nextRunner)
		go Runner(baston)
	}

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	if runner == 4 {
		fmt.Printf("Runner %d finish, Race owner\n", runner)
		wgChannelsTwo.Done()
		return
	}

	fmt.Printf("Runner %d exchange with runner %d\n", runner, nextRunner)

	baston <- nextRunner
}
