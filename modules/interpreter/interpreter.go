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
	matched, err := regexp.MatchString(`^"([^"]*)"`, value)
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

const (
	INTERPOLATION_SYMBOL_LEFT  string = "$("
	INTERPOLATION_SYMBOL_RIGHT string = ")"
)

func (interpreter *Interpreter) interpolateString(str string) string {
	interpolated := str
	for key, element := range interpreter.numberVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_SYMBOL_LEFT+key+INTERPOLATION_SYMBOL_RIGHT, fmt.Sprintf("%f", element))
	}
	for key, element := range interpreter.stringVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_SYMBOL_LEFT+key+INTERPOLATION_SYMBOL_RIGHT, element)
	}
	return interpolated
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
				interpreter.stringVarTable[variableName] = interpreter.interpolateString(value)
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
		case token.SUB:
			variableName := currentToken.GetParameter(0)
			subValue := currentToken.GetParameter(1)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Error: Referenced invalid variable '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(subValue); isRawNumber {
				interpreter.numberVarTable[variableName] -= value
			} else if interpreter.isNumberVar(subValue) {
				interpreter.numberVarTable[variableName] -= interpreter.numberVarTable[subValue]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", subValue, currentToken.GetLine())
			}
		case token.CONCAT:
			variableName := currentToken.GetParameter(0)
			concatValue := currentToken.GetParameter(1)
			if !interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Referenced invalid variable '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(concatValue); isRawNumber {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + fmt.Sprintf("%f", value)
			} else if isRawString, value := isRawString(concatValue); isRawString {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + interpreter.interpolateString(value)
			} else if interpreter.isNumberVar(concatValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + fmt.Sprintf("%f", interpreter.numberVarTable[concatValue])
			} else if interpreter.isStringVar(concatValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + interpreter.stringVarTable[concatValue]
			} else {
				log.Fatalf("Error: Referenced nonexistent variable '%s'. Line %d.", concatValue, currentToken.GetLine())
			}
		case token.SAY:
			output := currentToken.GetParameter(0)
			if isRawNumber, value := isRawNumber(output); isRawNumber {
				fmt.Println(value)
			} else if isRawString, value := isRawString(output); isRawString {
				fmt.Println(interpreter.interpolateString(value))
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
