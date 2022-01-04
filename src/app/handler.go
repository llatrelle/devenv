package app

import (
	"devEnv/src/container/command"
	"log"
)

func Exec(c *command.Command) {

	switch c.GetCommandName() {
	case "ps":
		err := command.PS(c.Args)
		handleError(err)
	case "logs":
		err := command.Logs(c.Args)
		handleError(err)
	case "image":
		err := command.PullImage(c.Args)
		handleError(err)
	case "start":
		err := command.Start(c.Args)
		handleError(err)
	case "create":
		err := command.Create(c.Args)
		handleError(err)
	case "stop":
		err := command.Stop(c.Args)
		handleError(err)
	case "serve":
		err := StartHTTP()
		handleError(err)
	default:
		log.Print("devEnv: invalid command")
	}
}

func handleError(err error) {
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}

}
