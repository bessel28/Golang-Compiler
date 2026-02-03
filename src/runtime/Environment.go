package runtime

import (
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	"errors"
	"fmt"
)

type RuntimeEnvironment struct {
	Variables map[string]*Literal
	Functions map[string]*FunctionSignature
	Parent    *RuntimeEnvironment
}

func (env *RuntimeEnvironment) Define(name string, value *Literal) error {
	if _, exists := env.Variables[name]; exists {
		return errors.New(fmt.Sprintf("Code Error: variable %v exists", name))
	}
	env.Variables[name] = value
	return nil
}

func (env *RuntimeEnvironment) Assign(name string, value *Literal) error {
	if _, exists := env.Variables[name]; !exists {
		return errors.New(fmt.Sprintf("Code Error: variable %v does not exist", name))
	}
	env.Variables[name] = value
	return nil
}
func (env *RuntimeEnvironment) Get(name string) (*Literal, error) {
	if _, exists := env.Variables[name]; !exists {
		if env.Parent != nil {
			return env.Parent.Get(name)
		}
		return nil, errors.New(fmt.Sprintf("Code Error: variable %v does not exist", name))
	}
	value := env.Variables[name]
	return value, nil
}

func (env *RuntimeEnvironment) NewChild() *RuntimeEnvironment {
	return &RuntimeEnvironment{
		Parent:    env,
		Variables: make(map[string]*Literal),
		Functions: make(map[string]*FunctionSignature),
	}
}
