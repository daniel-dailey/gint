package interpreter

import "fmt"

type SymbolType int

type SymbolIntf interface {
	str() string
	printname() string
}

const (
	SYMBOL_TYPE_NONE SymbolType = -1
)

type VarSymbol struct {
	*Symbol
	*BuiltInTypeSymbol
}

func (vs *VarSymbol) printname() string {
	return vs.name
}

func (vs *VarSymbol) str() string {
	return fmt.Sprintf("{%s:%s}", vs.name, vs.BuiltInTypeSymbol.str())
}

func InitVarSymbol(name string, sym *BuiltInTypeSymbol) *VarSymbol {
	return &VarSymbol{
		&Symbol{
			name: name,
		},
		sym,
	}
}

type BuiltInTypeSymbol struct {
	*Symbol
}

func (bi *BuiltInTypeSymbol) str() string {
	return bi.name
}

func (bi *BuiltInTypeSymbol) printname() string {
	return bi.name
}

func InitBuiltInTypeSymbol(name string) *BuiltInTypeSymbol {
	return &BuiltInTypeSymbol{
		&Symbol{
			name: name,
		},
	}
}

type Symbol struct {
	name string
	typ  SymbolType
}

func InitSymbol(name string, typ SymbolType) *Symbol {
	return &Symbol{
		name: name,
		typ:  typ,
	}
}

type SymbolTable struct {
	symbols map[string]SymbolIntf
}

func (st *SymbolTable) print() {
	for _, s := range st.symbols {
		fmt.Println(s.str())
	}
}

func (st *SymbolTable) insert(s SymbolIntf) {
	st.symbols[s.printname()] = s
}

func (st *SymbolTable) lookup(name string) SymbolIntf {
	return st.symbols[name]
}

func (st *SymbolTable) initBuiltinTypes() {
	st.insert(InitBuiltInTypeSymbol("INTEGER"))
	st.insert(InitBuiltInTypeSymbol("REAL"))
}

func InitSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: map[string]SymbolIntf{},
	}
}
