package command

// Command Database command interface
type Command interface {
	Cmd()
}

type Get struct {
	Cmd string
}
