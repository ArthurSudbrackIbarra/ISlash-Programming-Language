package interpreter

import (
	"fmt"
	"islash/modules/token"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func isFloat(value string) (bool, float64) {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false, -1
	}
	return true, parsedValue
}

func isString(value string) (bool, string) {
	matched, err := regexp.MatchString(`^"[a-zA-Z0-9!@#$&()\-.+,/ ]*"$`, value)
	if !matched || err != nil {
		return false, ""
	}
	return true, strings.ReplaceAll(value, "\"", "")
}

type Interpreter struct {
	variablesTable map[string]string
}

func (interpreter *Interpreter) isVar(value string) (bool, string) {
	if value, contains := interpreter.variablesTable[value]; contains {
		return true, value
	}
	return false, ""
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		variablesTable: make(map[string]string),
	}
}

func (interpreter *Interpreter) Interpret(tokensList []*token.Token) {
	for _, currentToken := range tokensList {
		switch currentToken.GetType() {
		case token.DECLARE:
			variableName := currentToken.GetParameter(0)
			assignValue := currentToken.GetParameter(1)
			isNumeric, _ := isFloat(assignValue)
			if !isNumeric {
				if isVar, value := interpreter.isVar(assignValue); isVar {
					interpreter.variablesTable[variableName] = value
				} else {
					log.Fatalf("Error: Referenced nonexistent variable '%s'.", assignValue)
				}
			} else {
				interpreter.variablesTable[variableName] = assignValue
			}
		case token.SAY:
			output := currentToken.GetParameter(0)
			if isString, value := isString(output); isString {
				fmt.Println(value)
			} else if isVar, value := interpreter.isVar(output); isVar {
				fmt.Println(value)
			} else {
				log.Fatalf("Error: Referenced nonexistent variable '%s'.", output)
			}
		}
	}
}
