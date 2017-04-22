package earth

import "github.com/stretchr/testify/mock"

type petWorldMock struct{ mock.Mock }

func (m *petWorldMock) Tallest(gs []Gopher) (Gopher, error) {
	args := m.Called(gs)
	return args.Get(0).(Gopher), args.Error(1)
}
