package interpreter

type ReturnType int

const (
	TYPE_FLOAT ReturnType = iota
	TYPE_INT
	TYPE_BOOL
	TYPE_NIL
)

type Type struct {
	token *Token
	value interface{}
}

func (t *Type) visit() (interface{}, ReturnType) {
	return nil, TYPE_NIL
}

func InitType(t *Token) *Type {
	return &Type{
		token: t,
		value: t.Value,
	}
}
