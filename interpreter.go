package main

import "fmt"

type Interpreter struct {
	parser *Parser
}

func (i *Interpreter) Interpret() {
	ret := i.parser.Parse()
	fmt.Println(ret.visit())
}

func NewInterpreter(p *Parser) *Interpreter {
	return &Interpreter{
		parser: p,
	}
}
