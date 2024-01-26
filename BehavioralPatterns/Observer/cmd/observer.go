package cmd

type Observer interface {
	update(string)
	getID() string
}
