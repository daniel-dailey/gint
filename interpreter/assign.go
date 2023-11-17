package interpreter

import (
	"log"
)

type Assign struct {
	left  AST
	right AST
	op    *Token
}

func (a *Assign) visit() (interface{}, ReturnType) {
	variableName := a.left.(*Var).val
	log.Println("assign visit: var name: ", variableName)
	v, rt := a.right.visit()
	if rt != TYPE_INT {
		log.Printf("ERROR: visiting assign... right node visit != type int??")
		return nil, TYPE_NIL
	}
	GLOBAL_SCOPE[variableName] = v.(int)
	log.Println("visited...")
	return nil, TYPE_NIL
}

func NewAssign(l, r AST, op *Token) *Assign {
	return &Assign{
		left:  l,
		right: r,
		op:    op,
	}
}
