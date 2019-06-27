package calc

import "errors"

type Calculator struct {
	result int
}

func (c *Calculator) Add(x int) {
	c.result += x
}

func (c *Calculator) Sub(x int) {
	c.result -= x
}

func (c *Calculator) MultiplyBy(x int) error {
	if x == 0 {
		return errors.New("zero not allowed")
	}
	c.result *= x
	return nil
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
