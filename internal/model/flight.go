package model

type Flight struct {
	name string
	cost float64
}

func NewFlight(name string, cost float64) *Flight {
	return &Flight{name, cost}
}

func (f *Flight) Name() (string, error) {
	return f.name, nil
}

func (f *Flight) Cost() (float64, error) {
	return f.cost, nil
}
