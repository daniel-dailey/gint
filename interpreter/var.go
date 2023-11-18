package interpreter

import (
	"strconv"
)

type Var struct {
	token *Token
	val   string
}

func (v *Var) visit() interface{} {
	variableName := v.val
	if val, ok := GLOBAL_SCOPE[variableName]; ok {
		return val
	}
	s, _ := strconv.Atoi(v.val)
	return s
}

func NewVar(t *Token) *Var {
	return &Var{
		token: t,
		val:   t.Value.(string),
	}
}
