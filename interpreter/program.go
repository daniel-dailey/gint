package interpreter

import "log"

type Program struct {
	name  string
	block AST
}

func (p *Program) visit() interface{} {
	log.Println("visit program...", p.name)
	p.block.visit()
	return nil
}

func InitProgram(name string, block AST) *Program {
	return &Program{
		name:  name,
		block: block,
	}
}
