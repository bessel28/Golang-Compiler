package runtime

import (
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	"errors"
	"fmt"
)

type Environment struct {
	variables map[string]*Literal
	functions map[string]*Function
	parent    *Environment
}

func (env *Environment) Define(name string, value *Literal) error {
	if _, exists := env.variables[name]; exists {
		return errors.New(fmt.Sprintf("Code Error: variable %v exists", name))
	}
	env.variables[name] = value
	return nil
}

func (env *Environment) Assign(name string, value *Literal) error {
	if _, exists := env.variables[name]; !exists {
		return errors.New(fmt.Sprintf("Code Error: variable %v does not exist", name))
	}
	env.variables[name] = value
	return nil
}
func (env *Environment) Get(name string) (*Literal, error) {
	if _, exists := env.variables[name]; !exists {
		if env.parent != nil {
			return env.parent.Get(name)
		}
		return nil, errors.New(fmt.Sprintf("Code Error: variable %v does not exist", name))
	}
	value := env.variables[name]
	return value, nil
}

func (env *Environment) NewChild() *Environment {
	return &Environment{
		parent:    env,
		variables: make(map[string]*Literal),
		functions: make(map[string]*Function),
	}
}
