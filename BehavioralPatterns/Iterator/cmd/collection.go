package cmd

type Collection interface {
	CreateIterator() Iterator
}
