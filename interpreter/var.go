package interpreter

import (
	"log"
	"strconv"
)

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
