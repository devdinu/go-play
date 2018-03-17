# Go Play
[![Build Status](https://travis-ci.org/devdinu/go-play.svg?branch=master)](https://travis-ci.org/devdinu/go-play)

 Samples of golang code and experiments, to explain concepts and to have fun.


### Content
 - go1.8 features, go1.10 features
 - go tools
 - pprof, tracing
 - channels
 - struct embedding
 - slice
 - [go testing](https://www.youtube.com/watch?v=zGhfJ88eKfw&index=1&list=PLKXvA3W4l9pHh2Pq04qCutB9e16QHMc26) (checkout branch gopherworld, and see git commits)

## Gotcha's

You think you've mastered Go in few weeks or days, well these are some interesting things or gotcha's you should be aware of.

- Method call on nil pointer is valid

```
type gopher struct {
	name string
}

func (g *gopher) Name() {
	fmt.Println("Gopher")
}

var g *gopher
fmt.Println("content: %v, address: %+p, output: %s", g, g, g.Name()) // content: <nil> address: +0x0, output: Gopher
```

- When you slice the original slice, and if you change the new sliced data, you would end up changing original slice. because the backing array for both of the slice is same.

```
 original := []int{1, 2, 3, 4, 5}
 new := original[0:3]
 new[0] = 1000
```

- Map is not addresseable, but slices are. Its because the map would grow when elements are added. so it would move the map content to different address whenever it grows.

```
	gs := map[string]gopher{}
	gs[0].name = "change"
    // Error: cannot assign to struct field gs["tall"].name in map
```

would change the original and have `{1000,2,3,4,5}`


## Concurrency
```
Do not communicate by sharing memory; instead, share memory by communicating
```

When we have multiple goroutines running, we 've to co-oridinate and share data with them to get the final outcome.


### Coordination

say we've increment function on a metric, if we've `n` goroutines we've to wait for it to complete before getting the final result for all of them to complete before getting the final call count.
```
// Incr(key string) 

for i := 0; i < n; i++ {
    go m.Incr("metrics.call")
}

m.Get("metrics.call") // would return inconsitent count i.e < n
```

we could do `time.Sleep(n * t)` if you know `t` is the processing time or naive way in tests. But The `Incr` function could increment the metric in cache, db or make network call, so mostly we can't define the time.

we can use `sync.WaitGroup` to coordinate. n goroutines so we say `wg.Add(n)` and each goroutine reports its completion with `wg.Done()`. Note that it have to be a pointer `*sync.WaitGroup`

```
var wg sync.WaitGroup
wg.Add(calls)

for i := 0; i < n; i++ {
	go func(wg *sync.WaitGroup) {
		m.Incr("incr.call")
		wg.Done()
	}(&wg)
}

wg.Wait() // waits till all n goroutines completes
```

### Communication

In the `Incr` since we've to maintain the metric count, we 've to store it. could be `map[string]int` to hold the count for each metric.
```
type metric struct {
    data map[string]int
}

...

func (m *metric) Incr(key string) {
	m.data[key] += 1
}
```

when goroutines concurrently access/change the map. we'll have inconsistent data, In our case map it will result in `concurrent map writes`. so we would 've to use locks andensure mutual exclusive access of data for writes. so the above incr would hold the lock and unlock after write.
```

type metric struct {
    data map[string]int
    sync.Mutex
}

...

func (m *metric) Incr(key string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.data[key] += 1
}
```

Acquiring locks would've its own cost, and if the implementation/logic becomes complex the time you would acquire a lock increase and will impact the performance.

Alternatively, we could share data across goroutines using channels.

```

type counter struct {
	occurence chan string
	data      map[string]int
}

func (c *counter) Incr(key string) {
	c.occurence <- key
}

func (c *counter) process() {
	for {
		ok, key := <-c.occurence
		if !ok { // Terimation case occurs, when chan is closed
			return
		}
		c.data[key] += 1
	}
}

go c.process() // a separate goroutine for receiving metric name, and increments the count.
```

In this [code](https://github.com/devdinu/go-play/blob/master/concurrency/metrics_chan.go#L18) we've used a single goroutine and achieved the same. No concurrent writes because there's only one goroutine which writes to the map. This is a contrived example, this case locks is more apt. but channels could be used in real world scenarios. workers, http throttling requests, collating responses from different HTTP apis etc.

Another simple example which uses channels to share data. Given a huge slice, main goroutine slices the original slice, and spins up goroutines and computes the sum of sub slice. 
The goroutines send the subsum via channel. we could've had the workers listen to a channel for `[]int` and send the sum in output channel instead of data[] int param.

```
// goroutine or worker which computes the sum
func add(data []int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, d := range data {
		sum += d
	}
	result <- sum // send the sum back to the collector
}
```

The collector sums up all the result of workers and returns the data with channel. 
```
func collect(subsum <-chan int, totalSum chan<- int) {
	total := 0
	for {
		s, ok := <-subsum
		if !ok {
			totalSum <- total
		}
		total += s
	}
}
```

check the complete [code](https://github.com/devdinu/go-play/blob/master/concurrency/pure_concurrency.go). 

Know more about buffered channels, how the goroutines waits while writing/reading, using `for { select { case` to wait without blocking, and closing channels and reading safely `data, ok := somechan`, and closing goroutines properly to avoid goroutine leaks before using it in prod.

[share memory by communicating](https://blog.golang.org/share-memory-by-communicating)


Look at gotchas code for more information.

## Slides of go meetup
- [go1.8 slides](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/go1.8.slide)
- [go tools slides](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/go_tools.slide)
- [Testing in go](https://youtu.be/zGhfJ88eKfw) Golang Meetup XXI
- [Embedding in go screencast](https://youtu.be/Ki3kUvEx4-8) Golang Meetup XXV
- [gotchas in go](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/gotchas.slide) Golang Meetup XXVI [screencast](https://youtu.be/J3plALnTjA8)
- [slice of Slices](https://goo.gl/NTmsqf) at [Golang Meetup](https://www.meetup.com/Golang-Bangalore/events/246437796/) XXVIII
- [go 1.10 slides](https://talks.godoc.org/github.com/devdinu/go-play/go1.10.release.slide#9) [screencast](https://youtu.be/t-iiICzV-es)
