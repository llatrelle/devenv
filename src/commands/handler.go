package commands

import "log"

func ExecuteCommand(c *Command) {

	switch c.GetCommandName() {
	case "ps":
		err := PSCommand(c.Args)
		handleError(err)
	case "logs":
		err := LogsCommand(c.Args)
		handleError(err)
	case "image":
		err := ImageCommand(c.Args)
		handleError(err)
	case "start":
		err := StartCommand(c.Args)
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
