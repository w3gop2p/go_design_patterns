package cmd

type CounterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

func (c *CounterTerroristDress) GetColor() string {
	return c.color
}
