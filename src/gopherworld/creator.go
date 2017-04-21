package gopherworld

type Gophers []Gopher

type gopherWorld interface {
	Tallest(gs Gophers) (Gopher, error)
}

type world struct{}

func (gw world) Tallest(gs Gophers) (Gopher, error) {
	return tallest(gs)
}

func Create() gopherWorld {
	return world{}
}
