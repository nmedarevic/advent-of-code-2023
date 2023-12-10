package workers

import (
	"advent_helper/channel"
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, id int, getCallback func(result *[]int) func(line string), lines <-chan []string, resultChanel chan<- []int) {
	defer wg.Done()

	fmt.Println("Worker", id, "started!")

	var result []int

	var callback = getCallback(&result)
	channel.IterateThroughInputChannel(lines, callback)

	fmt.Println("Worker", id, "finished!")

	resultChanel <- result
}

func WorkerInt(wg *sync.WaitGroup, id int, getCallback func(result *[]int) func(line int), intInput <-chan []int, resultChanel chan<- []int) {
	defer wg.Done()

	fmt.Println("Worker", id, "started!")

	var result []int

	var callback = getCallback(&result)
	channel.IterateThroughInputChannelInt(intInput, callback)

	fmt.Println("Worker", id, "finished!")

	resultChanel <- result
}

func MonitorWorker(wg *sync.WaitGroup, cs chan []int) {
	wg.Wait()
	close(cs)
}
