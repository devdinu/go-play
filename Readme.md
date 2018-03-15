# Go Play
[![Build Status](https://travis-ci.org/devdinu/go-play.svg?branch=master)](https://travis-ci.org/devdinu/go-play)

 Samples of golang code and experiments


## Content
 - Experimenting go1.8 features
 - go tools
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


Look at gotchas code for more information.

## Slides of go meetup
- [go1.8 slides](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/go1.8.slide)
- [go tools slides](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/go_tools.slide)
- [Testing in go](https://youtu.be/zGhfJ88eKfw) Golang Meetup XXI
- [Embedding in go screencast](https://youtu.be/Ki3kUvEx4-8) Golang Meetup XXV
- [gotchas in go](http://talks.godoc.org/github.com/dineshkumar-cse/go-play/gotchas.slide) Golang Meetup XXVI [screencast](https://youtu.be/J3plALnTjA8)
- [slice of Slices](https://goo.gl/NTmsqf) at [Golang Meetup](https://www.meetup.com/Golang-Bangalore/events/246437796/) XXVIII
- [go 1.10 slides](https://talks.godoc.org/github.com/devdinu/go-play/go1.10.release.slide#9) [screencast](https://youtu.be/t-iiICzV-es)
