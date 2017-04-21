package gopherworld

import "fmt"

type Gopher struct {
	name   string
	weight float64
	height int
}

func (g Gopher) String() string {
	return fmt.Sprintf("%s Gopher weighing: %f lb, %d mm tall\n", g.name, g.weight, g.height)
}
