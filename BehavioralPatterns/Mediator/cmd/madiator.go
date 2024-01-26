package cmd

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
