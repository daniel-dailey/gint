package interpreter

import (
	"strconv"
)

type Var struct {
	token *Token
	val   string
}

func (v *Var) visit() (interface{}, ReturnType) {
	variableName := v.val
	if val, ok := GLOBAL_SCOPE[variableName]; ok {
		return val, TYPE_INT
	}
	s, _ := strconv.Atoi(v.val)
	return s, TYPE_INT
}

func NewVar(t *Token) *Var {
	return &Var{
		token: t,
		val:   t.Value.(string),
	}
}
