package cmd

type server interface {
	HandleRequest(string, string) (int, string)
}
