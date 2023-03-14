package internal

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var (
	wgChannels sync.WaitGroup
)

func ChannelsOne() {
	runtime.GOMAXPROCS(1)
	// runtime.GOMAXPROCS(2) // Note a diferença da saida.

	channel := make(chan *string)
	wgChannels.Add(2)

	go runAt("A", channel)
	go runAt("B", channel)

	for n := range channel {
		fmt.Println(*n)
	}

	go func() {
		wgChannels.Wait()
		close(channel)
	}()

}

func runAt(id string, channel chan *string) {
	defer wgChannels.Done()

	for i := 0; i < 10; i++ {
		msg := id + ":" + strconv.Itoa(i)
		channel <- &msg
	}
}
