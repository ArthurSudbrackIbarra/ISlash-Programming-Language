package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	tokensList := lexer.MountTokens("example-programs/if-test.islash")
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
