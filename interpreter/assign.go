package interpreter

import (
	"log"
)

type Assign struct {
	left  AST
	right AST
	op    *Token
}

func (a *Assign) visit() interface{} {
	variableName := a.left.(*Var).val
	log.Println("assign visit: var name: ", variableName)
	v := a.right.visit()
	GLOBAL_SCOPE[variableName] = v
	return nil
}

func NewAssign(l, r AST, op *Token) *Assign {
	return &Assign{
		left:  l,
		right: r,
		op:    op,
	}
}
