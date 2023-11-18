package interpreter

type ReturnType int

const (
	TYPE_FLOAT ReturnType = iota
	TYPE_INT
	TYPE_BOOL
	TYPE_NIL
)

type Type struct {
	typ   TreeNodeType
	token *Token
	value interface{}
}

func (t *Type) val() interface{} {
	return t.value
}

func (t *Type) getType() TreeNodeType {
	return t.typ
}

func InitType(t *Token) *Type {
	return &Type{
		typ:   TreeNodeTypeType,
		token: t,
		value: t.Value,
	}
}
