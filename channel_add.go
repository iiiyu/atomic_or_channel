package main

func ChannelAdd() uint64 {
	one := uint64(1)
	numCh := make(chan uint64, 1)
	go func() {
		for l := 0; l < LOOP*CLIENTS; l++ {
			numCh <- one
		}
		close(numCh)
	}()
	return sum(numCh)
}

func sum(ch chan uint64) uint64 {
	var sum uint64
	for val := range ch {
		sum += val
	}
	return sum
}
