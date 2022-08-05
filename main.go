package main

import (
	"fmt"
	"islash/modules/lexer"
)

func main() {
	lexer := lexer.NewLexer("example-programs/ex-1.islash")
	tokensList := lexer.MountTokens()
	fmt.Printf("tokensList[0]: %v\n", tokensList[0])
}
