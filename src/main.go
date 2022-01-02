package main

import (
	"devEnv/src/app"
	"os"
)

func main() {
	args := os.Args[1:]
	app.Start(args)
}
