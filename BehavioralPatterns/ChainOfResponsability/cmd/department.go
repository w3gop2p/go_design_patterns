package cmd

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
