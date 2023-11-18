package interpreter

type Num struct {
	token *Token
	val   interface{}
}

func (n *Num) visit() interface{} {
	return n.val
}

func NewNum(t *Token) *Num {
	return &Num{
		token: t,
		val:   t.Value,
	}
}
