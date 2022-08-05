package interpreter

import (
	"islash/modules/token"
	"log"
	"strconv"
)

func isNumeric(value string) (bool, float64) {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false, -1
	}
	return true, parsedValue
}

type Interpreter struct {
	variablesTable map[string]string
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
			name := currentToken.GetParameter(0)
			value := currentToken.GetParameter(1)
			isNumeric, _ := isNumeric(value)
			if !isNumeric {
				if refVarValue, ok := interpreter.variablesTable[value]; ok {
					interpreter.variablesTable[name] = refVarValue
				} else {
					log.Fatalf("Error: Referenced nonexistent variable '%s'.", value)
				}
			} else {
				interpreter.variablesTable[name] = value
			}
		}
	}
}
