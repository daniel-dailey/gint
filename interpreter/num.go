package interpreter

type Num struct {
	token *Token
	val   interface{}
}

func (n *Num) visit() (interface{}, ReturnType) {
	switch n.val.(type) {
	case int:
		return n.val, TYPE_INT
	case float64:
		return n.val, TYPE_FLOAT
	default:
		return nil, TYPE_NIL
	}
	// return n.val, TYPE_INT
}

func NewNum(t *Token) *Num {
	return &Num{
		token: t,
		val:   t.Value,
	}
}
