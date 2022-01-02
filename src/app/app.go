package app

import (
	"devEnv/src/commands"
)

func Start(args []string) {

	c := new(commands.Command)
	c.ParseParams(args)
	commands.ExecuteCommand(c)

}
