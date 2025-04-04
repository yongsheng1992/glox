package core

import "fmt"

type Env struct {
	values map[string]interface{}
}

func NewEnv() *Env {
	return &Env{
		values: make(map[string]interface{}),
	}
}

func (env *Env) put(name string, value interface{}) {
	env.values[name] = value
}

func (env *Env) get(token *Token) interface{} {
	if value, exist := env.values[token.Lexeme]; !exist {
		panic(NewRuntimeError(token, fmt.Sprintf("Undefined variable '%s'.", token.Lexeme)))
	} else {
		return value
	}
}
