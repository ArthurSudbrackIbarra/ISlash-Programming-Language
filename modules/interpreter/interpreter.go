package interpreter

import (
	"bufio"
	"fmt"
	"islash/modules/token"
	"log"
	"math"
	"os"
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

func isRawNumberArray(value string) (bool, []float64) {
	return true, make([]float64, 0)
}

func isRawStringArray(value string) (bool, []string) {
	return true, make([]string, 0)
}

/*
	End - Helper functions
*/

type Interpreter struct {
	numberVarTable      map[string]float64
	stringVarTable      map[string]string
	numberArrayVarTable map[string][]float64
	stringArrayVarTable map[string][]string
	conditionStack      *Stack
	whileStack          *Stack
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		numberVarTable:      make(map[string]float64),
		stringVarTable:      make(map[string]string),
		numberArrayVarTable: make(map[string][]float64),
		stringArrayVarTable: make(map[string][]string),
		conditionStack:      NewEmptyStack(),
		whileStack:          NewEmptyStack(),
	}
}

func (interpreter *Interpreter) isNumberVar(key string) bool {
	if _, contains := interpreter.numberVarTable[key]; contains {
		return true
	}
	return false
}

func (interpreter *Interpreter) isStringVar(key string) bool {
	if _, contains := interpreter.stringVarTable[key]; contains {
		return true
	}
	return false
}

func (interpreter *Interpreter) isNumberArrayVar(key string) bool {
	if _, contains := interpreter.numberArrayVarTable[key]; contains {
		return true
	}
	return false
}

func (interpreter *Interpreter) isStringArrayVar(key string) bool {
	if _, contains := interpreter.stringArrayVarTable[key]; contains {
		return true
	}
	return false
}

func (interpreter *Interpreter) isVariableDefined(key string) bool {
	return interpreter.isNumberVar(key) || interpreter.isStringVar(key)
}

const (
	INTERPOLATION_LEFT  string = "$("
	INTERPOLATION_RIGHT string = ")"
)

func (interpreter *Interpreter) handleString(str string) string {
	interpolated := str
	for key, element := range interpreter.numberVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_LEFT+key+INTERPOLATION_RIGHT, strconv.FormatFloat(element, 'f', -1, 64))
	}
	for key, element := range interpreter.stringVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_LEFT+key+INTERPOLATION_RIGHT, element)
	}
	return strings.ReplaceAll(interpolated, `\n`, "\n")
}

func (interpreter *Interpreter) findMatchingEndWhileIndex(currentIndex int, tokensList []*token.Token) int {
	whileStatementsCount := 0
	currentWhileStatementOrder := 0
	for i := 0; i < len(tokensList); i++ {
		if tokensList[i].GetType() == token.WHILE {
			if i == currentIndex {
				currentWhileStatementOrder = whileStatementsCount
			}
			whileStatementsCount += 1
		}
	}
	endWhileToFindOrder := math.Abs(float64(currentWhileStatementOrder) - float64(whileStatementsCount))
	endwhilesFound := 0
	for i := 0; i < len(tokensList); i++ {
		if tokensList[i].GetType() == token.ENDWHILE {
			endwhilesFound += 1
			if endwhilesFound == int(endWhileToFindOrder) {
				return i
			}
		}
	}
	return -1
}

func (interpreter *Interpreter) findNextConditionBlockIndex(currentIndex int, tokensList []*token.Token) int {
	for i := currentIndex + 1; i < len(tokensList); i++ {
		if tokensList[i].GetType() == token.ENDIF {
			return i
		}
		if tokensList[i].GetType() == token.ELSE {
			return i
		}
	}
	return -1
}

func (interpreter *Interpreter) Interpret(tokensList []*token.Token) {
	for i := 0; i < len(tokensList); i++ {
		if i >= len(tokensList) {
			break
		}
		currentToken := tokensList[i]
		switch currentToken.GetType() {
		case token.DECLARE:
			variableName := currentToken.GetParameter(0)
			assignValue := currentToken.GetParameter(1)
			if isRawNumber, _ := isRawNumber(variableName); isRawNumber {
				log.Fatalf("Error: Variable name cannot be a number. Line %d.", currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(assignValue); isRawNumber {
				interpreter.numberVarTable[variableName] = value
			} else if isRawString, value := isRawString(assignValue); isRawString {
				interpreter.stringVarTable[variableName] = interpreter.handleString(value)
			} else if interpreter.isNumberVar(assignValue) {
				interpreter.numberVarTable[variableName] = interpreter.numberVarTable[assignValue]
			} else if interpreter.isStringVar(assignValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[assignValue]
			} else {
				log.Fatalf("Error: Invalid declaration of varible '%s'. Line %d.", variableName, currentToken.GetLine())
			}
		case token.ADD:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 + parsedAddValue2
		case token.SUB:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 - parsedAddValue2
		case token.MULT:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 * parsedAddValue2
		case token.DIV:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 / parsedAddValue2
		case token.MOD:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = math.Mod(parsedAddValue1, parsedAddValue2)
		case token.INCREMENT:
			variableName := currentToken.GetParameter(0)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', not a number variable. Line %d.", variableName, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] += 1
		case token.DECREMENT:
			variableName := currentToken.GetParameter(0)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', not a number variable. Line %d.", variableName, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] -= 1
		case token.GREATERTHAN:
			firstValue := currentToken.GetParameter(0)
			parsedFirstValue := -1.0
			secondValue := currentToken.GetParameter(1)
			parsedSecondValue := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(firstValue); isRawNumber {
				parsedFirstValue = value
			} else if interpreter.isNumberVar(firstValue) {
				parsedFirstValue = interpreter.numberVarTable[firstValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue); isRawNumber {
				parsedSecondValue = value
			} else if interpreter.isNumberVar(secondValue) {
				parsedSecondValue = interpreter.numberVarTable[secondValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if parsedFirstValue > parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.GREATERTHANEQUAL:
			firstValue := currentToken.GetParameter(0)
			parsedFirstValue := -1.0
			secondValue := currentToken.GetParameter(1)
			parsedSecondValue := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(firstValue); isRawNumber {
				parsedFirstValue = value
			} else if interpreter.isNumberVar(firstValue) {
				parsedFirstValue = interpreter.numberVarTable[firstValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue); isRawNumber {
				parsedSecondValue = value
			} else if interpreter.isNumberVar(secondValue) {
				parsedSecondValue = interpreter.numberVarTable[secondValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if parsedFirstValue >= parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.LESSTHAN:
			firstValue := currentToken.GetParameter(0)
			parsedFirstValue := -1.0
			secondValue := currentToken.GetParameter(1)
			parsedSecondValue := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(firstValue); isRawNumber {
				parsedFirstValue = value
			} else if interpreter.isNumberVar(firstValue) {
				parsedFirstValue = interpreter.numberVarTable[firstValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue); isRawNumber {
				parsedSecondValue = value
			} else if interpreter.isNumberVar(secondValue) {
				parsedSecondValue = interpreter.numberVarTable[secondValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if parsedFirstValue < parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.LESSTHANEQUAL:
			firstValue := currentToken.GetParameter(0)
			parsedFirstValue := -1.0
			secondValue := currentToken.GetParameter(1)
			parsedSecondValue := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(firstValue); isRawNumber {
				parsedFirstValue = value
			} else if interpreter.isNumberVar(firstValue) {
				parsedFirstValue = interpreter.numberVarTable[firstValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue); isRawNumber {
				parsedSecondValue = value
			} else if interpreter.isNumberVar(secondValue) {
				parsedSecondValue = interpreter.numberVarTable[secondValue]
			} else {
				log.Fatalf("Invalid parameter '%s', not a number or a number variable. Line %d.", firstValue, currentToken.GetLine())
			}
			if parsedFirstValue <= parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.EQUAL:
			firstValue := currentToken.GetParameter(0)
			secondValue := currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if interpreter.isNumberVar(firstValue) && interpreter.isStringVar(secondValue) {
				interpreter.numberVarTable[variableName] = 0
				return
			} else if interpreter.isNumberVar(secondValue) && interpreter.isStringVar(firstValue) {
				interpreter.numberVarTable[variableName] = 0
				return
			}
			isFirstValueRawNumber, numberValue1 := isRawNumber(firstValue)
			isFirstValueRawString, stringValue1 := isRawString(firstValue)
			isSecondValueRawNumber, numberValue2 := isRawNumber(secondValue)
			isSecondValueRawString, stringValue2 := isRawString(secondValue)
			if isFirstValueRawNumber && isSecondValueRawNumber {
				if numberValue1 == numberValue2 {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if isFirstValueRawString && isSecondValueRawString {
				if stringValue1 == stringValue2 {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if isFirstValueRawNumber && interpreter.isNumberVar(secondValue) {
				if numberValue1 == interpreter.numberVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if isSecondValueRawNumber && interpreter.isNumberVar(firstValue) {
				if numberValue2 == interpreter.numberVarTable[firstValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if isFirstValueRawString && interpreter.isStringVar(secondValue) {
				if stringValue1 == interpreter.stringVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if isSecondValueRawString && interpreter.isStringVar(firstValue) {
				if stringValue2 == interpreter.stringVarTable[firstValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if interpreter.isNumberVar(firstValue) && interpreter.isNumberVar(secondValue) {
				if interpreter.numberVarTable[firstValue] == interpreter.numberVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else if interpreter.isStringVar(firstValue) && interpreter.isStringVar(secondValue) {
				if interpreter.stringVarTable[firstValue] == interpreter.stringVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 1
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			} else {
				if !interpreter.isVariableDefined(firstValue) || !interpreter.isVariableDefined(secondValue) {
					log.Fatalf("Error: One or more variables are not defined. Line %d.", currentToken.GetLine())
				} else {
					interpreter.numberVarTable[variableName] = 0
				}
			}
		case token.NOTEQUAL:
			firstValue := currentToken.GetParameter(0)
			secondValue := currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if interpreter.isNumberVar(firstValue) && interpreter.isStringVar(secondValue) {
				interpreter.numberVarTable[variableName] = 1
				return
			} else if interpreter.isNumberVar(secondValue) && interpreter.isStringVar(firstValue) {
				interpreter.numberVarTable[variableName] = 1
				return
			}
			isFirstValueRawNumber, numberValue1 := isRawNumber(firstValue)
			isFirstValueRawString, stringValue1 := isRawString(firstValue)
			isSecondValueRawNumber, numberValue2 := isRawNumber(secondValue)
			isSecondValueRawString, stringValue2 := isRawString(secondValue)
			if isFirstValueRawNumber && isSecondValueRawNumber {
				if numberValue1 == numberValue2 {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if isFirstValueRawString && isSecondValueRawString {
				if stringValue1 == stringValue2 {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if isFirstValueRawNumber && interpreter.isNumberVar(secondValue) {
				if numberValue1 == interpreter.numberVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if isSecondValueRawNumber && interpreter.isNumberVar(firstValue) {
				if numberValue2 == interpreter.numberVarTable[firstValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if isFirstValueRawString && interpreter.isStringVar(secondValue) {
				if stringValue1 == interpreter.stringVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if isSecondValueRawString && interpreter.isStringVar(firstValue) {
				if stringValue2 == interpreter.stringVarTable[firstValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if interpreter.isNumberVar(firstValue) && interpreter.isNumberVar(secondValue) {
				if interpreter.numberVarTable[firstValue] == interpreter.numberVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if interpreter.isStringVar(firstValue) && interpreter.isStringVar(secondValue) {
				if interpreter.stringVarTable[firstValue] == interpreter.stringVarTable[secondValue] {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else {
				if !interpreter.isVariableDefined(firstValue) || !interpreter.isVariableDefined(secondValue) {
					log.Fatalf("Error: One or more variables are not defined. Line %d.", currentToken.GetLine())
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			}
		case token.NOT:
			notTarget := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(notTarget); isRawNumber {
				if value >= 1 {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else if interpreter.isNumberVar(notTarget) {
				if interpreter.numberVarTable[notTarget] >= 1 {
					interpreter.numberVarTable[variableName] = 0
				} else {
					interpreter.numberVarTable[variableName] = 1
				}
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", notTarget, currentToken.GetLine())
			}
		case token.AND:
			andTarget1 := currentToken.GetParameter(0)
			parsedTarget1 := -1.0
			andTarget2 := currentToken.GetParameter(1)
			parsedTarget2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(andTarget1); isRawNumber {
				parsedTarget1 = value
			} else if interpreter.isNumberVar(andTarget1) {
				parsedTarget1 = interpreter.numberVarTable[andTarget1]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", andTarget1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(andTarget2); isRawNumber {
				parsedTarget2 = value
			} else if interpreter.isNumberVar(andTarget2) {
				parsedTarget2 = interpreter.numberVarTable[andTarget2]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", andTarget2, currentToken.GetLine())
			}
			if parsedTarget1 >= 1 && parsedTarget2 >= 1 {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.OR:
			orTarget1 := currentToken.GetParameter(0)
			parsedTarget1 := -1.0
			orTarget2 := currentToken.GetParameter(1)
			parsedTarget2 := -1.0
			variableName := currentToken.GetParameter(2)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(orTarget1); isRawNumber {
				parsedTarget1 = value
			} else if interpreter.isNumberVar(orTarget1) {
				parsedTarget1 = interpreter.numberVarTable[orTarget1]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", orTarget1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(orTarget2); isRawNumber {
				parsedTarget2 = value
			} else if interpreter.isNumberVar(orTarget2) {
				parsedTarget2 = interpreter.numberVarTable[orTarget2]
			} else {
				log.Fatalf("Invalid parameter '%s'. Line %d.", orTarget2, currentToken.GetLine())
			}
			if parsedTarget1 >= 1 || parsedTarget2 >= 1 {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
		case token.IF:
			condition := currentToken.GetParameter(0)
			parsedCondition := -1.0
			if isRawNumber, value := isRawNumber(condition); isRawNumber {
				parsedCondition = value
			} else if isRawString, value := isRawString(condition); isRawString {
				if len(value) > 1 {
					parsedCondition = 1
				} else {
					parsedCondition = 0
				}
			} else if interpreter.isNumberVar(condition) {
				parsedCondition = interpreter.numberVarTable[condition]
			} else if interpreter.isStringVar(condition) {
				if len(interpreter.stringVarTable[condition]) > 1 {
					parsedCondition = 1
				} else {
					parsedCondition = 0
				}
			} else {
				log.Fatalf("Error: Invalid parameter for IF statement. Line %d.", currentToken.GetLine())
			}
			if parsedCondition < 1 {
				interpreter.conditionStack.Push(1)
				i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
				if i == -1 {
					log.Fatalf("Invalid IF block. Line %d", currentToken.GetLine())
				}
			} else {
				interpreter.conditionStack.Push(0)
			}
		case token.ELSE:
			shouldExecute := interpreter.conditionStack.Top()
			if shouldExecute != 1 {
				i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
				if i == -1 {
					log.Fatalf("Missing ENDIF statement. Line %d", currentToken.GetLine())
				}
			}
		case token.ENDIF:
			interpreter.conditionStack.Pop()
		case token.CONCAT:
			variableName := currentToken.GetParameter(0)
			concatValue := currentToken.GetParameter(1)
			if !interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Referenced invalid variable '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(concatValue); isRawNumber {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + fmt.Sprintf("%f", value)
			} else if isRawString, value := isRawString(concatValue); isRawString {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + interpreter.handleString(value)
			} else if interpreter.isNumberVar(concatValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + fmt.Sprintf("%f", interpreter.numberVarTable[concatValue])
			} else if interpreter.isStringVar(concatValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[variableName] + interpreter.stringVarTable[concatValue]
			} else {
				log.Fatalf("Error: Referenced nonexistent variable '%s'. Line %d.", concatValue, currentToken.GetLine())
			}
		case token.LENGTH:
			str := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			if interpreter.isStringVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a string variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawString, value := isRawString(str); isRawString {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.handleString(value)))
			} else if interpreter.isStringVar(str) {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.stringVarTable[str]))
			} else {
				log.Fatalf("Error: Invalid parameter '%s'. Line %d.", str, currentToken.GetLine())
			}
		case token.GETCHAR:
			str := currentToken.GetParameter(0)
			index := currentToken.GetParameter(1)
			parsedIndex := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(index); isRawNumber {
				parsedIndex = value
			} else if interpreter.isNumberVar(index) {
				parsedIndex = interpreter.numberVarTable[index]
			} else {
				log.Fatalf("Error: Invalid parameter '%s', must be a number or a number variable. Line %d.", index, currentToken.GetLine())
			}
			if interpreter.isNumberVar(variableName) {
				log.Fatalf("Error: Invalid parameter '%s', already a number variable. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawString, value := isRawString(str); isRawString {
				handledStr := interpreter.handleString(value)
				if int(parsedIndex) >= len(handledStr) {
					log.Fatalf("Error: Index out of bounds [%s] for string character. Line %d.", index, currentToken.GetLine())
				}
				interpreter.stringVarTable[variableName] = handledStr[int(parsedIndex) : int(parsedIndex)+1]
			} else if interpreter.isStringVar(str) {
				if int(parsedIndex) >= len(interpreter.stringVarTable[str]) {
					log.Fatalf("Error: Index out of bounds [%s] for string character. Line %d.", index, currentToken.GetLine())
				}
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[str][int(parsedIndex) : int(parsedIndex)+1]
			} else {
				log.Fatalf("Error: Invalid parameter '%s'. Line %d.", str, currentToken.GetLine())
			}
		case token.SAY:
			output := currentToken.GetParameter(0)
			if isRawNumber, value := isRawNumber(output); isRawNumber {
				fmt.Println(value)
			} else if isRawString, value := isRawString(output); isRawString {
				fmt.Println(interpreter.handleString(value))
			} else if interpreter.isNumberVar(output) {
				fmt.Println(interpreter.numberVarTable[output])
			} else if interpreter.isStringVar(output) {
				fmt.Println(interpreter.stringVarTable[output])
			} else {
				log.Fatalf("Error: Referenced nonexistent variable '%s'. Line %d.", output, currentToken.GetLine())
			}
		case token.INPUT:
			variableName := currentToken.GetParameter(0)
			text := currentToken.GetParameter(1)
			if isRawNumber, _ := isRawNumber(variableName); isRawNumber {
				log.Fatalf("Error: Invalid parameter '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawString, _ := isRawString(variableName); isRawString {
				log.Fatalf("Error: Invalid parameter '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawString, value := isRawString(text); isRawString {
				text = interpreter.handleString(value)
			}
			fmt.Print(text)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			if isRawNumber, value := isRawNumber(input); isRawNumber {
				interpreter.numberVarTable[variableName] = value
			} else {
				interpreter.stringVarTable[variableName] = input
			}
		case token.WHILE:
			condition := currentToken.GetParameter(0)
			if isRawNumber, value := isRawNumber(condition); isRawNumber {
				if value >= 1 {
					interpreter.whileStack.Push(i)
				} else {
					i = interpreter.findMatchingEndWhileIndex(i, tokensList)
					if i == -1 {
						log.Fatalf("Error: Missing ENDWHILE statement for WHILE in line %d.", currentToken.GetLine())
					}
				}
			} else if interpreter.isNumberVar(condition) {
				if interpreter.numberVarTable[condition] >= 1 {
					interpreter.whileStack.Push(i)
				} else {
					i = interpreter.findMatchingEndWhileIndex(i, tokensList)
					if i == -1 {
						log.Fatalf("Error: Missing ENDWHILE statement for WHILE in line %d.", currentToken.GetLine())
					}
				}
			}
		case token.ENDWHILE:
			indexToGoBack := interpreter.whileStack.Pop()
			if indexToGoBack != -1 {
				i = indexToGoBack - 1
			}
		}
	}
}
