package cmd

type Subject interface {
	Register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
