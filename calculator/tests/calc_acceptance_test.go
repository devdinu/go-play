package calc_test

import (
	"github.com/DATA-DOG/godog"
	calc "github.com/devdinu/go-play/calculator"
)

type CalcSuite struct {
	*godog.Suite
	calc calc.Calculator
}
