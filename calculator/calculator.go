package calc

type Calculator struct {
	result int
}

func (c *Calculator) Add(x int) {
	c.result += x
}

func (c *Calculator) Sub(x int) {
	c.result -= x
}

func (c *Calculator) Multiply(x int) {
	c.result *= x
}

func (c *Calculator) Press(x int) {
	c.result = x
}

func (c *Calculator) Clear() {
	c.result = 0
}

func (c *Calculator) Result() int {
	return c.result
}
