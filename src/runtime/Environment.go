package runtime

import (
	. "compiler/src/components/literals"
	. "compiler/src/components/statements"
	"errors"
	"fmt"
)

type RuntimeEnvironment struct {
	Variables map[string]*Literal
	Functions map[string]*FunctionStatement
	Parent    *RuntimeEnvironment
}

func (env *RuntimeEnvironment) DefineVariable(name string, value *Literal) error {
	if _, exists := env.Variables[name]; exists {
		return errors.New(fmt.Sprintf("Runtime Error: variable %v already exists", name))
	}
	env.Variables[name] = value
	return nil
}

func (env *RuntimeEnvironment) AssignVariable(name string, value *Literal) error {
	if _, exists := env.Variables[name]; !exists {
		return errors.New(fmt.Sprintf("Runtime Error: variable %v does not exist", name))
	}
	env.Variables[name] = value
	return nil
}
func (env *RuntimeEnvironment) GetVariable(name string) (*Literal, error) {
	if _, exists := env.Variables[name]; !exists {
		if env.Parent != nil {
			return env.Parent.GetVariable(name)
		}
		return nil, errors.New(fmt.Sprintf("Runtime Error: variable %v does not exist", name))
	}
	value := env.Variables[name]
	if value == nil {
		return nil, errors.New(fmt.Sprintf("Runtime Error: variable %v does not exist", name))
	}
	return value, nil
}

func (env *RuntimeEnvironment) DefineFunction(name string, function *FunctionStatement) error {
	if _, exists := env.Functions[name]; exists {
		return errors.New(fmt.Sprintf("Runtime Error: function %v already exists", name))
	}
	env.Functions[name] = function
	return nil
}

func (env *RuntimeEnvironment) GetFunction(name string) (*FunctionStatement, error) {
	if _, exists := env.Functions[name]; !exists {
		if env.Parent != nil {
			return env.Parent.GetFunction(name)
		}
		return nil, errors.New(fmt.Sprintf("Runtime Error: function %v does not exist", name))
	}
	function := env.Functions[name]
	if function == nil {
		return nil, errors.New(fmt.Sprintf("Runtime Error: function %v does not exist", name))
	}
	return function, nil
}

func (env *RuntimeEnvironment) NewChild() *RuntimeEnvironment {
	return &RuntimeEnvironment{
		Parent:    env,
		Variables: make(map[string]*Literal),
		Functions: make(map[string]*FunctionStatement),
	}
}
