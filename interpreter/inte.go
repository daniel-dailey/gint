package interpreter

import (
	"log"
)

type Interpreter struct {
	parser *Parser
}

var GLOBAL_SCOPE = map[string]int{}

func (i *Interpreter) Interpret() {
	ret := i.parser.Parse()
	ret.visit()
	log.Println(GLOBAL_SCOPE)
}

func NewInterpreter(p *Parser) *Interpreter {
	return &Interpreter{
		parser: p,
	}
}
