package model

import "fmt"

type Environment struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	BindPorts []string `json:"bindPorts"`
	Status    string   `json:"status"`
}

func (e *Environment) ToString() string {
	return fmt.Sprintf("\nname: %v\n"+
		"imagen: %v\n"+
		"status: %v\n"+
		"------", e.Name, e.Image, e.Status)
}
