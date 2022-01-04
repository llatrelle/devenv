package command

import (
	"strings"
)

type Param struct {
	Key   string
	Value string
}
type Params []Param

type Command struct {
	name string
	Args Params
}

func (c *Command) ParseParams(args []string) {
	if len(args) > 0 && !isKey(args[0]) {
		c.name = args[0]
	}

	for i, arg := range args {
		if isKey(arg) {
			param := Param{Key: strings.ToLower(arg)}
			if len(args)-1 > i && !isKey(args[i+1]) {

				param.Value = args[i+1]
			}
			//	log.Printf("len:%v i: %v isKey:%v", len(args), i, !isKey(args[i+1]))
			c.Args = append(c.Args, param)
		}
	}
}

//isKey valid if args is key
func isKey(k string) bool {
	return len(k) > 0 && k[0:1] == "-"
}

func (c *Command) GetCommandName() string {
	return c.name
}

func (p *Params) HaveArg(arg string) bool {
	lowerArg := strings.ToLower(arg)
	for _, a := range *p {
		if a.Key == lowerArg {
			return true
		}
	}
	return false
}

func (p *Params) GetValue(k string) string {
	for _, param := range *p {
		if k == param.Key {
			return param.Value
		}
	}
	return ""
}

func (p *Params) AddParam(key string) {
	*p = append(*p, Param{Key: key})

}
func (p *Params) AddParamWithValue(key, value string) {
	*p = append(*p, Param{Key: key, Value: value})
}