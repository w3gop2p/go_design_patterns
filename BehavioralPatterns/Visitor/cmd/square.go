package cmd

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}
