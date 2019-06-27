package calc_test

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	calc "github.com/devdinu/go-play/calculator"
)

type CalcSuite struct {
	*godog.Suite
	calc *calc.Calculator
}

func (cs *CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalcSuite) iPress(num int) error {
	cs.calc.Press(num)
	return nil
}

func (cs *CalcSuite) iAdd(num int) error {
	cs.calc.Add(num)
	return nil
}

func (cs *CalcSuite) iSubtract(num int) error {
	cs.calc.Sub(num)
	return nil
}

// Assertions
func (cs *CalcSuite) theResultShouldBe(expectedResult int) error {
	result := cs.calc.Result()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func (cs *CalcSuite) iMultiplyBy(factor int) error {
	return cs.calc.MultiplyBy(factor)
}

func FeatureContext(suite *godog.Suite) {
	s := CalcSuite{
		calc:  new(calc.Calculator),
		Suite: suite,
	}
	suite.Step(`^calculator is cleared$`, s.calculatorIsCleared)

	suite.Step(`^i press (\d+)$`, s.iPress)
	suite.Step(`^i add (\d+)$`, s.iAdd)
	suite.Step(`^i subtract (\d+)$`, s.iSubtract)
	suite.Step(`^i multiply by (\d+)$`, s.iMultiplyBy)

	suite.Step(`^the result should be (\d+)$`, s.theResultShouldBe)
}
