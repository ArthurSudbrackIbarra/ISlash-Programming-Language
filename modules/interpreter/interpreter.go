package interpreter

import (
	"fmt"
	"islash/modules/token"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
	Start - Helper functions
*/

func isRawNumber(value string) (bool, float64) {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false, -1
	}
	return true, parsedValue
}

func isRawString(value string) (bool, string) {
	matched, err := regexp.MatchString(`^"[a-zA-Z0-9!@#$&()\-.+,/ ]*"$`, value)
	if !matched || err != nil {
		return false, ""
	}
	return true, strings.ReplaceAll(value, "\"", "")
}

/*
	End - Helper functions
*/

type Interpreter struct {
	numberVarTable map[string]float64
	stringVarTable map[string]string
}

func (interpreter *Interpreter) isNumberVar(value string) bool {
	if _, contains := interpreter.numberVarTable[value]; contains {
		return true
	}
	return false
}

func (interpreter *Interpreter) isStringVar(value string) bool {
	if _, contains := interpreter.stringVarTable[value]; contains {
		return true
	}
	return false
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		numberVarTable: make(map[string]float64),
		stringVarTable: make(map[string]string),
	}
}

func (interpreter *Interpreter) Interpret(tokensList []*token.Token) {
	for i := 0; i < len(tokensList); i++ {
		currentToken := tokensList[i]
		switch currentToken.GetType() {
		case token.DECLARE:
			variableName := currentToken.GetParameter(0)
			assignValue := currentToken.GetParameter(1)
			if isRawNumber, value := isRawNumber(assignValue); isRawNumber {
				interpreter.numberVarTable[variableName] = value
			} else if isRawString, value := isRawString(assignValue); isRawString {
				interpreter.stringVarTable[variableName] = value
			} else if interpreter.isNumberVar(assignValue) {
				interpreter.numberVarTable[variableName] = interpreter.numberVarTable[assignValue]
			} else if interpreter.isStringVar(assignValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[assignValue]
			} else {
				log.Fatalf("Invalid declaration of varible '%s'. Line %d.", variableName, currentToken.GetLine())
			}
		case token.ADD:
			variableName := currentToken.GetParameter(0)
			addValue := currentToken.GetParameter(1)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Error: Referenced invalid variable '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue); isRawNumber {
				interpreter.numberVarTable[variableName] += value
			} else if interpreter.isNumberVar(addValue) {
				interpreter.numberVarTable[variableName] += interpreter.numberVarTable[addValue]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", addValue, currentToken.GetLine())
			}
		case token.SAY:
			output := currentToken.GetParameter(0)
			if isRawNumber, value := isRawNumber(output); isRawNumber {
				fmt.Println(value)
			} else if isRawString, value := isRawString(output); isRawString {
				fmt.Println(value)
			} else if interpreter.isNumberVar(output) {
				fmt.Println(interpreter.numberVarTable[output])
			} else if interpreter.isStringVar(output) {
				fmt.Println(interpreter.stringVarTable[output])
			} else {
				log.Fatalf("Error: Referenced nonexistent variable '%s'. Line %d.", output, currentToken.GetLine())
			}
		}
	}
}
