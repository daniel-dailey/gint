package interpreter

import (
	"fmt"
)

type Interpreter struct {
	parser *Parser
}

var GLOBAL_SCOPE = map[string]interface{}{}

func (i *Interpreter) Interpret() {
	rootNode := i.parser.Parse()
	rootNode.visit()
	fmt.Println("======================================================")
	fmt.Println("=                      OUTPUT                        =")
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println("=                    MEMORY MAP                      =")
	fmt.Println("======================================================")
	for k, v := range GLOBAL_SCOPE {
		fmt.Printf(">> %s = %d\n", k, v)
	}
	fmt.Println("======================================================")
}

func NewInterpreter(p *Parser) *Interpreter {
	return &Interpreter{
		parser: p,
	}
}
