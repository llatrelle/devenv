package app

import (
	"devEnv/src/container/command"
)

func Init(args []string) {
	c := new(command.Command)
	c.ParseParams(args)
	Exec(c)

}

func Serve() {

}
