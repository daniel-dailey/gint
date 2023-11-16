package main

type Num struct {
	token *Token
	val   int
}

func (n *Num) visit() int {
	return n.val
}

func NewNum(t *Token) *Num {
	return &Num{
		token: t,
		val:   t.Value,
	}
}
