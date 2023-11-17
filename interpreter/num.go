package interpreter

import "log"

type Num struct {
	token *Token
	val   interface{}
}

func (n *Num) visit() (interface{}, ReturnType) {
	log.Println("Num visit: ", n.val)
	return n.val, TYPE_INT
}

func NewNum(t *Token) *Num {
	log.Println("NEW NUM: ", t.String(), t.Value)
	return &Num{
		token: t,
		val:   t.Value,
	}
}
