package interpreter

import (
	"log"
	"strconv"
)

type Compound struct {
	children []AST
}

func (c *Compound) visit() (interface{}, ReturnType) {
	for _, child := range c.children {
		child.visit()
	}
	return nil, TYPE_NIL
}

func NewCompound() *Compound {
	return &Compound{
		children: make([]AST, 0),
	}
}

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

type Var struct {
	token *Token
	val   string
}

func (v *Var) visit() (interface{}, ReturnType) {
	variableName := v.val
	log.Println("visit var: var name = ", variableName)
	if val, ok := GLOBAL_SCOPE[variableName]; ok {
		return val, TYPE_INT
	}
	s, _ := strconv.Atoi(v.val)
	// log.Fatalln("name err... ", variableName, v.token.Value)
	return s, TYPE_INT
}

func NewVar(t *Token) *Var {
	return &Var{
		token: t,
		val:   t.Value.(string),
	}
}
