package eprunner

import (
	"github.com/nonemax/eprunner/process"
	"github.com/nonemax/eprunner/runner"
)

// New return new process struct
func New(name, args string) process.Data {
	re := runner.New()
	p := process.Data{
		Name: name,
		Args: args,
		Exec: &re,
	}
	return p
}
