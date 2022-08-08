package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
)

func main() {
	/*
		if len(os.Args) < 2 {
			log.Fatal("No program name specified.")
		}
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Unnable to get user's current directory.")
		}
		sourceCodePath := filepath.Join(cwd, os.Args[1])
	*/
	sourceCodePath := "programs/length.islash"
	tokensList := lexer.MountTokens(sourceCodePath)
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList)
}
