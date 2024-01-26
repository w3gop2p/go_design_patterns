package cmd

type Device interface {
	on()
	off()
}
