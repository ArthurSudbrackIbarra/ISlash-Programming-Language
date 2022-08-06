package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	tokensList := lexer.MountTokens("example-programs/ex-1.islash")
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
