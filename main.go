package main

import (
	"islash/modules/interpreter"
	"islash/modules/lexer"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	/*
		Finding the program path.
	*/

	if len(os.Args) < 2 {
		log.Fatal("No program name specified.")
	}
	if !strings.HasSuffix(os.Args[1], ".isl") {
		log.Fatal("Your program extension must be '.isl'.")
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unnable to get user's current directory.")
	}
	sourceCodePath := filepath.Join(cwd, os.Args[1])

	/*
		Mounting tokens with the lexer.
	*/
	tokensList := lexer.MountTokens(sourceCodePath)
	/*
		Interpreting the tokens with the interpreter.
	*/
	sourceCodeDir := filepath.Join(sourceCodePath, "..")
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(tokensList, sourceCodeDir)
}
