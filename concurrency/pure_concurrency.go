package concurrency

import (
	"fmt"
	"sync"
)

func Add(data []int, split int) int {
	start := 0
	end := 0
	sum := make(chan int, 1)
	var wg sync.WaitGroup
	for end < len(data) {
		end = start + split
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		// instantiating separate gorouting to sum sliced data
		go add(data[start:end], sum, &wg)
		start += split
	}

	totalSum := make(chan int)
	go collect(sum, totalSum)

	wg.Wait()
	close(sum)

	s := <-totalSum
	fmt.Println("Final Sum:", s)
	return s
}

func collect(subsum <-chan int, totalSum chan<- int) {
	total := 0
	for s := range subsum {
		total += s
	}
	totalSum <- total
}

func add(data []int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, d := range data {
		sum += d
	}
	result <- sum
}

func SumIt(n int, concurrency int) int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}
	return Add(data, concurrency)
}
