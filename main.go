package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	tokensList := lexer.MountTokens("example-programs/is-prime-number.islash")
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
