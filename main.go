package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	tokensList := lexer.MountTokens("example-programs/sum-of-first-x-numbers.islash")
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
