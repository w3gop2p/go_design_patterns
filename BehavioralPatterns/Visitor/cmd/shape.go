package cmd

type Shape interface {
	getType() string
	accept(Visitor)
}
