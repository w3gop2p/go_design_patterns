package cmd

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
