package cmd

type Train interface {
	Arrive()
	Depart()
	permitArrival()
}
