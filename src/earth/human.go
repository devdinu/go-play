package earth

type Gopher struct {
	name   string
	weight float64
	Height int
}

var totalPets int

type human struct {
	PetWorld
	name string
	pets []Gopher
}

func (h human) tallestPet() (Gopher, error) {
	return h.Tallest(h.pets)
	//return Gopher{}, nil
}

type PetWorld interface {
	Tallest([]Gopher) (Gopher, error)
}

func (h human) addPet(g Gopher) {
	totalPets += 1
}

func (h human) hatePets() {
	totalPets = 0
}

func (h human) race() {
	go h.addPet(Gopher{})
	go h.hatePets()
}
