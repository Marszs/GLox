package interpreter

import (
	le "GLox/internal/loxerror"
	"GLox/internal/scanner/token"
)

// Environment 用来管理变量名->值之间的映射

type Environment struct {
	enclosing *Environment
	values    map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing: enclosing, values: make(map[string]interface{})}
}

func (e *Environment) define(name *token.Token, value interface{}) {
	e.values[name.Lexeme] = value
}

func (e *Environment) defineLiteral(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) lookup(name *token.Token) (interface{}, error) {
	// 当存在作用域的时候，现在当前作用域查找
	if value, exist := e.values[name.Lexeme]; exist {
		return value, nil
	}
	// 再往上层作用域（enclosing）查找，上层会（递归）查找上层的上层...直到遍历完所有作用域
	if e.enclosing != nil {
		return e.enclosing.lookup(name)
	}

	//panic(le.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'."))
	return nil, le.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}

func (e *Environment) assign(name *token.Token, value interface{}) error {
	if _, exist := e.values[name.Lexeme]; exist {
		e.values[name.Lexeme] = value
		return nil
	}
	// 如果当前作用域不存在赋值的变量，则再往上层作用域寻找并赋值
	// var a = 1;
	// {
	//	 a = 2
	// }
	if e.enclosing != nil {
		return e.enclosing.assign(name, value)
	}

	//panic(le.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'."))
	return le.NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}

func (e *Environment) ancestor(distance int) *Environment {
	environment := e
	for i := 0; i < distance; i++ {
		environment = environment.enclosing
	}

	return environment
}

func (e *Environment) getAt(distance int, name string) interface{} {
	return e.ancestor(distance).values[name]
}

func (e *Environment) assignAt(distance int, token *token.Token, value interface{}) {
	e.ancestor(distance).values[token.Lexeme] = value
}
