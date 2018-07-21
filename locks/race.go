package locks

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	val int
}

func (c *Counter) Incr() {
	c.val++
}

func doJobs(as, ss int) int {
	cntr := new(Counter)

	var wg sync.WaitGroup
	wg.Add(as + ss)
	ticker := time.NewTicker(time.Millisecond * 400)
	go func() {
		for t := range ticker.C {
			fmt.Printf("[%v] Counter: %d\n", t, cntr.val)
		}
	}()

	go Update(cntr, 1, as, &wg)

	go Update(cntr, -1, ss, &wg)

	wg.Wait()
	fmt.Println("done....")
	ticker.Stop()
	fmt.Println(cntr.val)
	return cntr.val
}

func Update(cntr *Counter, delta, totalJobs int, wg *sync.WaitGroup) {
	for i := 0; i < totalJobs; i++ {
		go func() {
			defer wg.Done()
			//some Operation
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			cntr.val += delta
		}()
	}
}
