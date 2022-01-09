package controller

import (
	"devEnv/src/api/common"
	"devEnv/src/container/command"
	"net/http"
)

func GetEnvironments(w http.ResponseWriter, r *http.Request) {
	p := new(command.Params)
	p.AddParam("-all")
	envs, err := command.PS(*p)
	if err != nil {
		common.Success(w, err)
	}
	common.Success(w, envs)

}
