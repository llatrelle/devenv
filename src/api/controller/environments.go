package controller

import (
	"devEnv/src/container/command"
	"net/http"
)

func GetEnvironments(w http.ResponseWriter, r *http.Request) {
	p := new(command.Params)
	p.AddParam("-all")
	command.PS(*p)
	w.Write([]byte("asd"))
}
