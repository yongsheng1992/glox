package core

import "fmt"

type Env struct {
	values    map[string]interface{}
	enclosing *Env
}

func NewEnv() *Env {
	return &Env{
		values: make(map[string]interface{}),
	}
}

func NewEnvWithEnclosing(env *Env) *Env {
	return &Env{
		values:    make(map[string]interface{}),
		enclosing: env,
	}
}

func (env *Env) assign(token *Token, value interface{}) {
	if _, exist := env.values[token.Lexeme]; exist {
		env.values[token.Lexeme] = value
		return
	}
	if env.enclosing != nil {
		env.enclosing.assign(token, value)
		return
	}

	panic(NewRuntimeError(token, fmt.Sprintf("Undefined variable '%s'.", token.Lexeme)))
}

func (env *Env) define(token *Token, value interface{}) {
	env.values[token.Lexeme] = value
}

func (env *Env) get(token *Token) interface{} {
	if value, exist := env.values[token.Lexeme]; !exist {
		panic(NewRuntimeError(token, fmt.Sprintf("Undefined variable '%s'.", token.Lexeme)))
	} else {
		return value
	}
}
