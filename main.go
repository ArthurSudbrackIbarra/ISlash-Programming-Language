package main

import (
	"fmt"
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	tokensList := lexer.MountTokens("example-programs/ex-1.islash")
	fmt.Printf("tokensList[0]: %v\n", tokensList[0])
	fmt.Printf("tokensList[1]: %v\n", tokensList[1])
	fmt.Printf("tokensList[2]: %v\n", tokensList[2])
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
