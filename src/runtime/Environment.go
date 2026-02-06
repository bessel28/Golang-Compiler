package runtime

import (
	. "compiler/src/components/statements"
	. "compiler/src/runtime/value"
	"errors"
	"fmt"
)

type RuntimeEnvironment struct {
	Variables map[string]RuntimeValue
	Functions map[FunctionIdentifier]*FunctionStatement
	Parent    *RuntimeEnvironment
}

func NewRuntimeEnvironment(vars map[string]RuntimeValue, funcs map[FunctionIdentifier]*FunctionStatement, parent *RuntimeEnvironment) *RuntimeEnvironment {
	if vars == nil {
		vars = make(map[string]RuntimeValue)
	}
	if funcs == nil {
		funcs = make(map[FunctionIdentifier]*FunctionStatement)
	}

	return &RuntimeEnvironment{vars, funcs, parent}
}

func (env *RuntimeEnvironment) DefineVariable(name string, value RuntimeValue) error {
	if _, exists := env.Variables[name]; exists {
		return errors.New(fmt.Sprintf("Runtime Error: variable %v already exists", name))
	}
	env.Variables[name] = value
	return nil
}

func (env *RuntimeEnvironment) AssignVariable(name string, value RuntimeValue) error {
	if _, exists := env.Variables[name]; !exists {
		return errors.New(fmt.Sprintf("Runtime Error: variable %v does not exist", name))
	}
	env.Variables[name] = value
	return nil
}
func (env *RuntimeEnvironment) GetVariable(name string) (RuntimeValue, error) {
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

func (env *RuntimeEnvironment) DefineFunction(identifier *FunctionIdentifier, function *FunctionStatement) error {
	if _, exists := env.Functions[*identifier]; exists {
		return errors.New(fmt.Sprintf("Runtime Error: function %v already exists", identifier.String()))
	}
	env.Functions[*identifier] = function
	return nil
}

func (env *RuntimeEnvironment) GetFunction(identifier *FunctionIdentifier) (*FunctionStatement, error) {
	if _, exists := env.Functions[*identifier]; !exists {
		if env.Parent != nil {
			return env.Parent.GetFunction(identifier)
		}
		return nil, errors.New(fmt.Sprintf("Runtime Error: function %v does not exist", identifier.Name))
	}
	function := env.Functions[*identifier]
	if function == nil {
		return nil, errors.New(fmt.Sprintf("Runtime Error: function %v does not exist", identifier.Name))
	}
	return function, nil
}

func (env *RuntimeEnvironment) NewChild() *RuntimeEnvironment {
	return NewRuntimeEnvironment(nil, nil, env)
}
