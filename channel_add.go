package main

import (
	"runtime"
	"sync/atomic"
)

func ChannelAdd() uint64 {
	runtime.GOMAXPROCS(0)
	var count uint64

	// start worker
	sum := func(ch chan *uint64) {
		for value := range ch {
			atomic.AddUint64(&count, *value)
		}
	}

	one := uint64(1)
	for i := 1; i <= CLIENTS; i++ {
		ch := make(chan *uint64, LOOP)
		go func() {
			for l := 0; l < LOOP; l++ {
				ch <- &one
			}
			close(ch)
		}()
		sum(ch)
	}

	return count

}
