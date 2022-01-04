package model

type Environment struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	BindPorts []string `json:"bindPorts"`
}
