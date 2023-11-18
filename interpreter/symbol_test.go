package interpreter

import (
	"log"
	"testing"
)

func TestBasicTypes(t *testing.T) {
	intType := InitBuiltInTypeSymbol("INTEGER")
	realType := InitBuiltInTypeSymbol("REAL")
	varXSymbol := InitVarSymbol("x", intType)
	varYSymbol := InitVarSymbol("y", realType)
	log.Println(varXSymbol.str(), varYSymbol.str())
}

func TestSymbolTable(t *testing.T) {
	st := InitSymbolTable()
	intType := InitBuiltInTypeSymbol("INTEGER")
	st.insert(intType)
	varXSymbol := InitVarSymbol("x", intType)
	st.insert(varXSymbol)

	realType := InitBuiltInTypeSymbol("REAL")
	varYSymbol := InitVarSymbol("y", realType)
	st.insert(varYSymbol)
	st.print()
}
