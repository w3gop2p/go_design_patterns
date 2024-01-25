package cmd

type TerroristDress struct {
	color string
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

func (t *TerroristDress) GetColor() string {
	return t.color
}
