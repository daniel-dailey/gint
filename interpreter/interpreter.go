package interpreter

import (
	"log"
)

type Interpreter struct {
	parser *Parser
}

var GLOBAL_SCOPE = map[string]int{}

func (i *Interpreter) Interpret() {
	rootNode := i.parser.Parse()
	rootNode.visit()
	log.Println(GLOBAL_SCOPE)
	for k, v := range GLOBAL_SCOPE {
		log.Printf("%s = %d", k, v)
	}
}

func NewInterpreter(p *Parser) *Interpreter {
	return &Interpreter{
		parser: p,
	}
}
