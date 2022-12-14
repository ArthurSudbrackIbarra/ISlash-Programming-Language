package interpreter

import (
	"bufio"
	"fmt"
	"islash/modules/io"
	"islash/modules/token"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func containsNumberArray(array []float64, element float64) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
}

func containsStrArray(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
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
	foreachIndexesStack *Stack
	foreachNamesStack   *Stack
	lastForeachEnded    bool
	varsToDelete        *Stack
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		numberVarTable:      make(map[string]float64),
		stringVarTable:      make(map[string]string),
		numberArrayVarTable: make(map[string][]float64),
		stringArrayVarTable: make(map[string][]string),
		conditionStack:      NewEmptyStack(),
		whileStack:          NewEmptyStack(),
		foreachIndexesStack: NewEmptyStack(),
		foreachNamesStack:   NewEmptyStack(),
		lastForeachEnded:    false,
		varsToDelete:        NewEmptyStack(),
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

func (interpreter *Interpreter) isRawNumberArray(value string) (bool, []float64) {
	numberArray := make([]float64, 0)
	if value == "[]number" {
		return true, numberArray
	}
	if !strings.HasPrefix(value, "[") || !strings.HasSuffix(value, "]") {
		return false, nil
	}
	value = value[1 : len(value)-1]
	splittedStr := strings.Split(value, ",")
	for _, element := range splittedStr {
		if isRawNumber, number := isRawNumber(element); isRawNumber {
			numberArray = append(numberArray, number)
		} else if interpreter.isNumberVar(element) {
			numberArray = append(numberArray, interpreter.numberVarTable[element])
		} else {
			return false, nil
		}
	}
	return true, numberArray
}

func (interpreter *Interpreter) isRawStringArray(value string) (bool, []string) {
	strArray := make([]string, 0)
	if value == "[]string" {
		return true, strArray
	}
	if !strings.HasPrefix(value, "[") || !strings.HasSuffix(value, "]") {
		return false, nil
	}
	value = value[1 : len(value)-1]
	splittedStr := strings.Split(value, ",")
	for _, element := range splittedStr {
		if isRawString, str := isRawString(element); isRawString {
			strArray = append(strArray, str)
		} else if interpreter.isStringVar(element) {
			strArray = append(strArray, interpreter.stringVarTable[element])
		} else {
			return false, nil
		}
	}
	return true, strArray
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
	for key, element := range interpreter.numberArrayVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_LEFT+key+INTERPOLATION_RIGHT, strings.Join(strings.Fields(fmt.Sprint(element)), " "))
	}
	for key, element := range interpreter.stringArrayVarTable {
		interpolated = strings.ReplaceAll(interpolated, INTERPOLATION_LEFT+key+INTERPOLATION_RIGHT, strings.Join(strings.Fields(fmt.Sprint(element)), " "))
	}
	return strings.ReplaceAll(interpolated, `\n`, "\n")
}

func (interpreter *Interpreter) findCloseLoopIndex(currentIndex int, tokensList []*token.Token) int {
	stack := NewEmptyStack()
	for i := 0; i < len(tokensList); i++ {
		if tokensList[i].GetType() == token.WHILE {
			stack.Push(i)
		} else if tokensList[i].GetType() == token.ENDWHILE {
			if i >= currentIndex {
				return i
			} else {
				stack.Pop()
			}
		}
	}
	return -1
}

func (interpreter *Interpreter) findNextConditionBlockIndex(currentIndex int, tokensList []*token.Token) int {
	levels := [][]int{}
	currentLevel := -1
	targetLevel := -1
	for i := 0; i < len(tokensList); i++ {
		if tokensList[i].GetType() == token.IF {
			if currentLevel == len(levels)-1 {
				levels = append(levels, []int{i})
				currentLevel += 1
			} else {
				currentLevel += 1
				if currentLevel >= len(levels[currentLevel]) {
					return -1
				}
				levels[currentLevel] = append(levels[currentLevel], i)
			}
		} else if tokensList[i].GetType() == token.ELSEIF || tokensList[i].GetType() == token.ELSE {
			levels[currentLevel] = append(levels[currentLevel], i)
		} else if tokensList[i].GetType() == token.ENDIF {
			levels[currentLevel] = append(levels[currentLevel], i)
			currentLevel -= 1
		}
		if i == currentIndex {
			targetLevel = currentLevel
		}
	}
	if targetLevel == -1 {
		return -1
	}
	for i := 0; i < len(levels[targetLevel]); i++ {
		conditionIndex := levels[targetLevel][i]
		if conditionIndex == currentIndex {
			if i+1 < len(levels[targetLevel]) {
				return levels[targetLevel][i+1]
			} else {
				return -1
			}
		}
	}
	return -1
}

func (interpreter *Interpreter) deleteVarIfSameName(varName string, varType string) {
	for key := range interpreter.numberVarTable {
		if key == varName && varType != "number" {
			delete(interpreter.numberVarTable, key)
		}
	}
	for key := range interpreter.stringVarTable {
		if key == varName && varType != "string" {
			delete(interpreter.stringVarTable, key)
		}
	}
	for key := range interpreter.numberArrayVarTable {
		if key == varName && varType != "numberarray" {
			delete(interpreter.numberArrayVarTable, key)
		}
	}
	for key := range interpreter.stringArrayVarTable {
		if key == varName && varType != "stringarray" {
			delete(interpreter.stringArrayVarTable, key)
		}
	}
}

func (interpreter *Interpreter) deleteVars() {
	for {
		if interpreter.varsToDelete.IsEmpty() {
			break
		}
		varName := interpreter.varsToDelete.Pop().(string)
		delete(interpreter.numberVarTable, varName)
		delete(interpreter.stringVarTable, varName)
		delete(interpreter.numberArrayVarTable, varName)
		delete(interpreter.stringArrayVarTable, varName)
	}
}

func (interpreter *Interpreter) Interpret(tokensList []*token.Token, sourceCodeDir string) {
	for i := 0; i < len(tokensList); i++ {
		if i >= len(tokensList) {
			break
		}
		currentToken := tokensList[i]
		switch currentToken.GetType() {
		case token.VAR:
			variableName := currentToken.GetParameter(0)
			assignValue := currentToken.GetParameter(1)
			if isRawNumber, _ := isRawNumber(variableName); isRawNumber {
				log.Fatalf("Interpreter error: Variable name cannot be a number. Line %d.", currentToken.GetLine())
			} else if isRawString, _ := isRawString(variableName); isRawString {
				log.Fatalf("Interpreter error: Variable name cannot be a string. Line %d.", currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(assignValue); isRawNumber {
				interpreter.numberVarTable[variableName] = value
				interpreter.deleteVarIfSameName(variableName, "number")
			} else if isRawString, value := isRawString(assignValue); isRawString {
				interpreter.stringVarTable[variableName] = interpreter.handleString(value)
				interpreter.deleteVarIfSameName(variableName, "string")
			} else if interpreter.isNumberVar(assignValue) {
				interpreter.numberVarTable[variableName] = interpreter.numberVarTable[assignValue]
				interpreter.deleteVarIfSameName(variableName, "number")
			} else if interpreter.isStringVar(assignValue) {
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[assignValue]
				interpreter.deleteVarIfSameName(variableName, "string")
			} else if isRawNumberArray, value := interpreter.isRawNumberArray(assignValue); isRawNumberArray {
				interpreter.numberArrayVarTable[variableName] = value
				interpreter.deleteVarIfSameName(variableName, "numberarray")
			} else if isRawStringArray, value := interpreter.isRawStringArray(assignValue); isRawStringArray {
				interpreter.stringArrayVarTable[variableName] = value
				interpreter.deleteVarIfSameName(variableName, "stringarray")
			} else if interpreter.isNumberArrayVar(assignValue) {
				interpreter.numberArrayVarTable[variableName] = interpreter.numberArrayVarTable[assignValue]
				interpreter.deleteVarIfSameName(variableName, "numberarray")
			} else if interpreter.isStringArrayVar(assignValue) {
				interpreter.stringArrayVarTable[variableName] = interpreter.stringArrayVarTable[assignValue]
				interpreter.deleteVarIfSameName(variableName, "stringarray")
			} else {
				log.Fatalf("Interpreter error: Invalid declaration of varible '%s'. Line %d.", variableName, currentToken.GetLine())
			}
		case token.ADD:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 + parsedAddValue2
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.SUB:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 - parsedAddValue2
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.MULT:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 * parsedAddValue2
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.DIV:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = parsedAddValue1 / parsedAddValue2
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.MOD:
			addValue1 := currentToken.GetParameter(0)
			parsedAddValue1 := -1.0
			addValue2 := currentToken.GetParameter(1)
			parsedAddValue2 := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(addValue1); isRawNumber {
				parsedAddValue1 = value
			} else if interpreter.isNumberVar(addValue1) {
				parsedAddValue1 = interpreter.numberVarTable[addValue1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(addValue2); isRawNumber {
				parsedAddValue2 = value
			} else if interpreter.isNumberVar(addValue2) {
				parsedAddValue2 = interpreter.numberVarTable[addValue2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", addValue2, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = math.Mod(parsedAddValue1, parsedAddValue2)
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.POWER:
			powerNumber := currentToken.GetParameter(0)
			parsedPowerNumber := -1.0
			powerValue := currentToken.GetParameter(1)
			parsedPowerValue := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(powerNumber); isRawNumber {
				parsedPowerNumber = value
			} else if interpreter.isNumberVar(powerNumber) {
				parsedPowerNumber = interpreter.numberVarTable[powerNumber]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", powerNumber, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(powerValue); isRawNumber {
				parsedPowerValue = value
			} else if interpreter.isNumberVar(powerValue) {
				parsedPowerValue = interpreter.numberVarTable[powerValue]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", powerValue, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = math.Pow(parsedPowerNumber, parsedPowerValue)
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.ROOT:
			rootNumber := currentToken.GetParameter(0)
			parsedPowerRoot := -1.0
			rootValue := currentToken.GetParameter(1)
			parsedRootValue := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(rootNumber); isRawNumber {
				parsedPowerRoot = value
			} else if interpreter.isNumberVar(rootNumber) {
				parsedPowerRoot = interpreter.numberVarTable[rootNumber]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", rootNumber, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(rootValue); isRawNumber {
				parsedRootValue = value
			} else if interpreter.isNumberVar(rootValue) {
				parsedRootValue = interpreter.numberVarTable[rootValue]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", rootValue, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] = math.Pow(parsedPowerRoot, 1/parsedRootValue)
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.INCREMENT:
			variableName := currentToken.GetParameter(0)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number variable. Line %d.", variableName, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] += 1
		case token.DECREMENT:
			variableName := currentToken.GetParameter(0)
			if !interpreter.isNumberVar(variableName) {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number variable. Line %d.", variableName, currentToken.GetLine())
			}
			interpreter.numberVarTable[variableName] -= 1
		case token.GREATER:
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
			if parsedFirstValue > parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.GREATEREQUAL:
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
			if parsedFirstValue >= parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.LESS:
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
			if parsedFirstValue < parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.LESSEQUAL:
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
			if parsedFirstValue <= parsedSecondValue {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.EQUAL:
			var firstValue interface{}
			firstValue = currentToken.GetParameter(0)
			var secondValue interface{}
			secondValue = currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(firstValue.(string)); isRawNumber {
				firstValue = value
			} else if interpreter.isNumberVar(firstValue.(string)) {
				firstValue = interpreter.numberVarTable[firstValue.(string)]
			} else if isRawString, value := isRawString(firstValue.(string)); isRawString {
				firstValue = value
			} else if interpreter.isStringVar(firstValue.(string)) {
				firstValue = interpreter.stringVarTable[firstValue.(string)]
			} else if isNumberArrayVar, value := interpreter.isRawNumberArray(firstValue.(string)); isNumberArrayVar {
				firstValue = value
			} else if interpreter.isNumberArrayVar(firstValue.(string)) {
				firstValue = interpreter.numberArrayVarTable[firstValue.(string)]
			} else if isStringArrayVar, value := interpreter.isRawStringArray(firstValue.(string)); isStringArrayVar {
				firstValue = value
			} else if interpreter.isStringArrayVar(firstValue.(string)) {
				firstValue = interpreter.stringArrayVarTable[firstValue.(string)]
			} else {
				log.Fatalf("Error Invalid parameter '%s', variable not defined. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue.(string)); isRawNumber {
				secondValue = value
			} else if interpreter.isNumberVar(secondValue.(string)) {
				secondValue = interpreter.numberVarTable[secondValue.(string)]
			} else if isRawString, value := isRawString(secondValue.(string)); isRawString {
				secondValue = value
			} else if interpreter.isStringVar(secondValue.(string)) {
				secondValue = interpreter.stringVarTable[secondValue.(string)]
			} else if isNumberArrayVar, value := interpreter.isRawNumberArray(secondValue.(string)); isNumberArrayVar {
				secondValue = value
			} else if interpreter.isNumberArrayVar(secondValue.(string)) {
				secondValue = interpreter.numberArrayVarTable[secondValue.(string)]
			} else if isStringArrayVar, value := interpreter.isRawStringArray(secondValue.(string)); isStringArrayVar {
				secondValue = value
			} else if interpreter.isStringArrayVar(secondValue.(string)) {
				secondValue = interpreter.stringArrayVarTable[secondValue.(string)]
			} else {
				log.Fatalf("Error Invalid parameter '%s', variable not defined. Line %d.", secondValue, currentToken.GetLine())
			}
			if reflect.DeepEqual(firstValue, secondValue) {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.NOTEQUAL:
			var firstValue interface{}
			firstValue = currentToken.GetParameter(0)
			var secondValue interface{}
			secondValue = currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(firstValue.(string)); isRawNumber {
				firstValue = value
			} else if interpreter.isNumberVar(firstValue.(string)) {
				firstValue = interpreter.numberVarTable[firstValue.(string)]
			} else if isRawString, value := isRawString(firstValue.(string)); isRawString {
				firstValue = value
			} else if interpreter.isStringVar(firstValue.(string)) {
				firstValue = interpreter.stringVarTable[firstValue.(string)]
			} else if isNumberArrayVar, value := interpreter.isRawNumberArray(firstValue.(string)); isNumberArrayVar {
				firstValue = value
			} else if interpreter.isNumberArrayVar(firstValue.(string)) {
				firstValue = interpreter.numberArrayVarTable[firstValue.(string)]
			} else if isStringArrayVar, value := interpreter.isRawStringArray(firstValue.(string)); isStringArrayVar {
				firstValue = value
			} else if interpreter.isStringArrayVar(firstValue.(string)) {
				firstValue = interpreter.stringArrayVarTable[firstValue.(string)]
			} else {
				log.Fatalf("Error Invalid parameter '%s', variable not defined. Line %d.", firstValue, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(secondValue.(string)); isRawNumber {
				secondValue = value
			} else if interpreter.isNumberVar(secondValue.(string)) {
				secondValue = interpreter.numberVarTable[secondValue.(string)]
			} else if isRawString, value := isRawString(secondValue.(string)); isRawString {
				secondValue = value
			} else if interpreter.isStringVar(secondValue.(string)) {
				secondValue = interpreter.stringVarTable[secondValue.(string)]
			} else if isNumberArrayVar, value := interpreter.isRawNumberArray(secondValue.(string)); isNumberArrayVar {
				secondValue = value
			} else if interpreter.isNumberArrayVar(secondValue.(string)) {
				secondValue = interpreter.numberArrayVarTable[secondValue.(string)]
			} else if isStringArrayVar, value := interpreter.isRawStringArray(secondValue.(string)); isStringArrayVar {
				secondValue = value
			} else if interpreter.isStringArrayVar(secondValue.(string)) {
				secondValue = interpreter.stringArrayVarTable[secondValue.(string)]
			} else {
				log.Fatalf("Error Invalid parameter '%s', variable not defined. Line %d.", secondValue, currentToken.GetLine())
			}
			if !reflect.DeepEqual(firstValue, secondValue) {
				interpreter.numberVarTable[variableName] = 1
			} else {
				interpreter.numberVarTable[variableName] = 0
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.NOT:
			notTarget := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
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
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.AND:
			andTarget1 := currentToken.GetParameter(0)
			parsedTarget1 := -1.0
			andTarget2 := currentToken.GetParameter(1)
			parsedTarget2 := -1.0
			variableName := currentToken.GetParameter(2)
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
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.OR:
			orTarget1 := currentToken.GetParameter(0)
			parsedTarget1 := -1.0
			orTarget2 := currentToken.GetParameter(1)
			parsedTarget2 := -1.0
			variableName := currentToken.GetParameter(2)
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
			interpreter.deleteVarIfSameName(variableName, "number")
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
				log.Fatalf("Interpreter error: Invalid parameter for IF statement. Line %d.", currentToken.GetLine())
			}
			if parsedCondition < 1 {
				interpreter.conditionStack.Push(1)
				i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
				if i == -1 {
					log.Fatalf("Interpreter error: Invalid IF block. Line %d", currentToken.GetLine())
				}
			} else {
				interpreter.conditionStack.Push(0)
			}
		case token.ELSEIF:
			shouldExecute := interpreter.conditionStack.Pop()
			if shouldExecute == 1 {
				condition := currentToken.GetParameter(0)
				parsedCondition := -1.0
				if isRawNumber, value := isRawNumber(condition); isRawNumber {
					parsedCondition = value
				} else if interpreter.isNumberVar(condition) {
					parsedCondition = interpreter.numberVarTable[condition]
				} else if interpreter.isStringVar(condition) {
					if len(interpreter.stringVarTable[condition]) > 1 {
						parsedCondition = 1
					} else {
						parsedCondition = 0
					}
				} else if isRawString, value := isRawString(condition); isRawString {
					if len(value) > 1 {
						parsedCondition = 1
					} else {
						parsedCondition = 0
					}
				} else {
					log.Fatalf("Interpreter error: Invalid parameter for ELSEIF statement. Line %d.", currentToken.GetLine())
				}
				if parsedCondition < 1 {
					interpreter.conditionStack.Push(1)
					i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
					if i == -1 {
						log.Fatalf("Interpreter error: Invalid IF block. Line %d", currentToken.GetLine())
					}
				} else {
					interpreter.conditionStack.Push(0)
				}
			} else if shouldExecute == 0 {
				interpreter.conditionStack.Push(0)
				i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
				if i == -1 {
					log.Fatalf("Interpreter error: Invalid IF block. Line %d", currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid IF block. Line %d", currentToken.GetLine())
			}
		case token.ELSE:
			shouldExecute := interpreter.conditionStack.Top()
			if shouldExecute == 0 {
				i = interpreter.findNextConditionBlockIndex(i, tokensList) - 1
				if i == -1 {
					log.Fatalf("Interpreter error: Missing ENDIF statement. Line %d", currentToken.GetLine())
				}
			} else if shouldExecute != 1 {
				log.Fatalf("Interpreter error: Invalid IF block. Line %d", currentToken.GetLine())
			}
		case token.ENDIF:
			interpreter.conditionStack.Pop()
		case token.CONCAT:
			variableName := currentToken.GetParameter(0)
			concatValue := currentToken.GetParameter(1)
			if !interpreter.isStringVar(variableName) {
				log.Fatalf("Interpreter error: Referenced invalid variable '%s'. Line %d.", variableName, currentToken.GetLine())
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
				log.Fatalf("Interpreter error: Referenced nonexistent variable '%s'. Line %d.", concatValue, currentToken.GetLine())
			}
		case token.UPPER:
			str := currentToken.GetParameter(0)
			parsedStr := ""
			variableName := currentToken.GetParameter(1)
			if isRawString, value := isRawString(str); isRawString {
				parsedStr = interpreter.handleString(value)
			} else if interpreter.isStringVar(str) {
				parsedStr = interpreter.stringVarTable[str]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a raw string or string variable. Line %d.", str, currentToken.GetLine())
			}
			interpreter.stringVarTable[variableName] = strings.ToUpper(parsedStr)
			interpreter.deleteVarIfSameName(variableName, "string")
		case token.LOWER:
			str := currentToken.GetParameter(0)
			parsedStr := ""
			variableName := currentToken.GetParameter(1)
			if isRawString, value := isRawString(str); isRawString {
				parsedStr = interpreter.handleString(value)
			} else if interpreter.isStringVar(str) {
				parsedStr = interpreter.stringVarTable[str]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a raw string or string variable. Line %d.", str, currentToken.GetLine())
			}
			interpreter.stringVarTable[variableName] = strings.ToLower(parsedStr)
			interpreter.deleteVarIfSameName(variableName, "string")
		case token.CONTAINS:
			var target interface{}
			target = currentToken.GetParameter(0)
			var toFind interface{}
			toFind = currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			if isRawString, value := isRawString(target.(string)); isRawString {
				target = interpreter.handleString(value)
			} else if interpreter.isStringVar(target.(string)) {
				target = interpreter.stringVarTable[target.(string)]
			} else if isRawNumberArray, value := interpreter.isRawNumberArray(target.(string)); isRawNumberArray {
				target = value
			} else if interpreter.isNumberArrayVar(target.(string)) {
				target = interpreter.numberArrayVarTable[target.(string)]
			} else if isRawStringArray, value := interpreter.isRawStringArray(target.(string)); isRawStringArray {
				target = value
			} else if interpreter.isStringArrayVar(target.(string)) {
				target = interpreter.stringArrayVarTable[target.(string)]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a string or an array. Line %d.", target, currentToken.GetLine())
			}
			if isRawString, value := isRawString(toFind.(string)); isRawString {
				toFind = interpreter.handleString(value)
			} else if interpreter.isStringVar(toFind.(string)) {
				toFind = interpreter.stringVarTable[toFind.(string)]
			} else if isRawNumber, value := isRawNumber(toFind.(string)); isRawNumber {
				toFind = value
			} else if interpreter.isNumberVar(toFind.(string)) {
				toFind = interpreter.numberVarTable[toFind.(string)]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a string or a number. Line %d.", toFind, currentToken.GetLine())
			}
			if reflect.TypeOf(target).String() == "string" {
				if reflect.TypeOf(toFind).String() == "string" {
					if strings.Contains(target.(string), toFind.(string)) {
						interpreter.numberVarTable[variableName] = 1
					} else {
						interpreter.numberVarTable[variableName] = 0
					}
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a string. Line %d.", toFind, currentToken.GetLine())
				}
			} else if reflect.TypeOf(target).String() == "[]float64" {
				if reflect.TypeOf(toFind).String() == "float64" {
					if containsNumberArray(target.([]float64), toFind.(float64)) {
						interpreter.numberVarTable[variableName] = 1
					} else {
						interpreter.numberVarTable[variableName] = 0
					}
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a number. Line %d.", toFind, currentToken.GetLine())
				}
			} else if reflect.TypeOf(target).String() == "[]string" {
				if reflect.TypeOf(toFind).String() == "string" {
					if containsStrArray(target.([]string), toFind.(string)) {
						interpreter.numberVarTable[variableName] = 1
					} else {
						interpreter.numberVarTable[variableName] = 0
					}
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a string. Line %d.", toFind, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a string or an array. Line %d.", target, currentToken.GetLine())
			}
		case token.LENGTH:
			target := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			if isRawString, value := isRawString(target); isRawString {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.handleString(value)))
			} else if interpreter.isStringVar(target) {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.stringVarTable[target]))
			} else if isRawNumberArray, value := interpreter.isRawNumberArray(target); isRawNumberArray {
				interpreter.numberVarTable[variableName] = float64(len(value))
			} else if isRawStringArray, value := interpreter.isRawStringArray(target); isRawStringArray {
				interpreter.numberVarTable[variableName] = float64(len(value))
			} else if interpreter.isNumberArrayVar(target) {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.numberArrayVarTable[target]))
			} else if interpreter.isStringArrayVar(target) {
				interpreter.numberVarTable[variableName] = float64(len(interpreter.stringArrayVarTable[target]))
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", target, currentToken.GetLine())
			}
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.CHARAT:
			str := currentToken.GetParameter(0)
			index := currentToken.GetParameter(1)
			parsedIndex := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(index); isRawNumber {
				parsedIndex = value
			} else if interpreter.isNumberVar(index) {
				parsedIndex = interpreter.numberVarTable[index]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', must be a number or a number variable. Line %d.", index, currentToken.GetLine())
			}
			if isRawString, value := isRawString(str); isRawString {
				handledStr := interpreter.handleString(value)
				if int(parsedIndex) >= len(handledStr) {
					log.Fatalf("Interpreter error: Index out of bounds [%s] for string character. Line %d.", index, currentToken.GetLine())
				}
				interpreter.stringVarTable[variableName] = handledStr[int(parsedIndex) : int(parsedIndex)+1]
			} else if interpreter.isStringVar(str) {
				if int(parsedIndex) >= len(interpreter.stringVarTable[str]) {
					log.Fatalf("Interpreter error: Index out of bounds [%s] for string character. Line %d.", index, currentToken.GetLine())
				}
				interpreter.stringVarTable[variableName] = interpreter.stringVarTable[str][int(parsedIndex) : int(parsedIndex)+1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", str, currentToken.GetLine())
			}
			interpreter.deleteVarIfSameName(variableName, "string")
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
			} else if isRawNumberArray, value := interpreter.isRawNumberArray(output); isRawNumberArray {
				fmt.Println(value)
			} else if isRawStringArray, value := interpreter.isRawStringArray(output); isRawStringArray {
				fmt.Println(value)
			} else if interpreter.isNumberArrayVar(output) {
				fmt.Println(interpreter.numberArrayVarTable[output])
			} else if interpreter.isStringArrayVar(output) {
				fmt.Println(interpreter.stringArrayVarTable[output])
			} else {
				log.Fatalf("Interpreter error: Referenced nonexistent variable '%s'. Line %d.", output, currentToken.GetLine())
			}
		case token.INPUT:
			variableName := currentToken.GetParameter(0)
			text := currentToken.GetParameter(1)
			if isRawNumber, _ := isRawNumber(variableName); isRawNumber {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", variableName, currentToken.GetLine())
			}
			if isRawString, _ := isRawString(variableName); isRawString {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", variableName, currentToken.GetLine())
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
					i = interpreter.findCloseLoopIndex(i, tokensList)
					if i == -1 {
						log.Fatalf("Interpreter error: Missing ENDWHILE statement for WHILE in line %d.", currentToken.GetLine())
					}
				}
			} else if interpreter.isNumberVar(condition) {
				if interpreter.numberVarTable[condition] >= 1 {
					interpreter.whileStack.Push(i)
				} else {
					i = interpreter.findCloseLoopIndex(i, tokensList)
					if i == -1 {
						log.Fatalf("Interpreter error: Missing ENDWHILE statement for WHILE in line %d.", currentToken.GetLine())
					}
				}
			}
		case token.BREAK:
			if interpreter.whileStack.IsEmpty() {
				log.Fatalf("Interpreter error: Cannot use BREAK instruction here. %d.", currentToken.GetLine())
			}
			interpreter.whileStack.Pop()
			i = interpreter.findCloseLoopIndex(i, tokensList)
			if i == -1 {
				log.Fatalf("Interpreter error: Cannot use BREAK instruction here. %d.", currentToken.GetLine())
			}
		case token.ENDWHILE:
			indexToGoBack := interpreter.whileStack.Pop()
			if indexToGoBack != -1 {
				i = indexToGoBack.(int) - 1
			}
		case token.APPEND:
			array := currentToken.GetParameter(0)
			element := currentToken.GetParameter(1)
			if interpreter.isNumberArrayVar(array) {
				if isRawNumber, value := isRawNumber(element); isRawNumber {
					interpreter.numberArrayVarTable[array] = append(interpreter.numberArrayVarTable[array], value)
				} else if interpreter.isNumberVar(element) {
					interpreter.numberArrayVarTable[array] = append(interpreter.numberArrayVarTable[array], interpreter.numberVarTable[element])
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a number. Line %d.", element, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				if isRawString, value := isRawString(element); isRawString {
					interpreter.stringArrayVarTable[array] = append(interpreter.stringArrayVarTable[array], value)
				} else if interpreter.isStringVar(element) {
					interpreter.stringArrayVarTable[array] = append(interpreter.stringArrayVarTable[array], interpreter.stringVarTable[element])
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a string. Line %d.", element, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.PREPEND:
			array := currentToken.GetParameter(0)
			element := currentToken.GetParameter(1)
			if interpreter.isNumberArrayVar(array) {
				if isRawNumber, value := isRawNumber(element); isRawNumber {
					interpreter.numberArrayVarTable[array] = append([]float64{value}, interpreter.numberArrayVarTable[array]...)
				} else if interpreter.isNumberVar(element) {
					interpreter.numberArrayVarTable[array] = append([]float64{interpreter.numberVarTable[element]}, interpreter.numberArrayVarTable[array]...)
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a number. Line %d.", element, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				if isRawString, value := isRawString(element); isRawString {
					interpreter.stringArrayVarTable[array] = append([]string{value}, interpreter.stringArrayVarTable[array]...)
				} else if interpreter.isStringVar(element) {
					interpreter.stringArrayVarTable[array] = append([]string{interpreter.stringVarTable[element]}, interpreter.stringArrayVarTable[array]...)
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a string. Line %d.", element, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.REMOVELAST:
			array := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			if interpreter.isNumberArrayVar(array) {
				if len(interpreter.numberArrayVarTable[array]) > 0 {
					lastElement := interpreter.numberArrayVarTable[array][len(interpreter.numberArrayVarTable[array])-1]
					interpreter.numberArrayVarTable[array] = interpreter.numberArrayVarTable[array][:len(interpreter.numberArrayVarTable[array])-1]
					interpreter.numberVarTable[variableName] = lastElement
					interpreter.deleteVarIfSameName(variableName, "number")
				} else {
					log.Fatalf("Interpreter error: Array '%s' is empty. Line %d.", array, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				if len(interpreter.stringArrayVarTable[array]) > 0 {
					lastElement := interpreter.stringArrayVarTable[array][len(interpreter.stringArrayVarTable[array])-1]
					interpreter.stringArrayVarTable[array] = interpreter.stringArrayVarTable[array][:len(interpreter.stringArrayVarTable[array])-1]
					interpreter.stringVarTable[variableName] = lastElement
					interpreter.deleteVarIfSameName(variableName, "string")
				} else {
					log.Fatalf("Interpreter error: Array '%s' is empty. Line %d.", array, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.REMOVEFIRST:
			array := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			if interpreter.isNumberArrayVar(array) {
				if len(interpreter.numberArrayVarTable[array]) > 0 {
					firstElement := interpreter.numberArrayVarTable[array][0]
					interpreter.numberArrayVarTable[array] = interpreter.numberArrayVarTable[array][1:]
					interpreter.numberVarTable[variableName] = firstElement
					interpreter.deleteVarIfSameName(variableName, "number")
				} else {
					log.Fatalf("Interpreter error: Array '%s' is empty. Line %d.", array, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				if len(interpreter.stringArrayVarTable[array]) > 0 {
					firstElement := interpreter.stringArrayVarTable[array][0]
					interpreter.stringArrayVarTable[array] = interpreter.stringArrayVarTable[array][1:]
					interpreter.stringVarTable[variableName] = firstElement
					interpreter.deleteVarIfSameName(variableName, "string")
				} else {
					log.Fatalf("Interpreter error: Array '%s' is empty. Line %d.", array, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.SETINDEX:
			array := currentToken.GetParameter(0)
			index := currentToken.GetParameter(1)
			parsedIndex := -1.0
			newValue := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(index); isRawNumber {
				parsedIndex = value
			} else if interpreter.isNumberVar(index) {
				parsedIndex = interpreter.numberVarTable[index]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", index, currentToken.GetLine())
			}
			if interpreter.isNumberArrayVar(array) {
				parsedNewValue := -1.0
				if isRawNumber, value := isRawNumber(newValue); isRawNumber {
					parsedNewValue = value
				} else if interpreter.isNumberVar(newValue) {
					parsedNewValue = interpreter.numberVarTable[newValue]
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", newValue, currentToken.GetLine())
				}
				if parsedIndex >= 0 && parsedIndex < float64(len(interpreter.numberArrayVarTable[array])) {
					interpreter.numberArrayVarTable[array][int(parsedIndex)] = parsedNewValue
				} else {
					log.Fatalf("Interpreter error: Index '%f' out of bounds. Line %d.", parsedIndex, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				parsedNewValue := ""
				if isRawString, value := isRawString(newValue); isRawString {
					parsedNewValue = value
				} else if interpreter.isStringVar(newValue) {
					parsedNewValue = interpreter.stringVarTable[newValue]
				} else {
					log.Fatalf("Interpreter error: Invalid parameter '%s', not a string or a string variable. Line %d.", newValue, currentToken.GetLine())
				}
				if parsedIndex >= 0 && parsedIndex < float64(len(interpreter.stringArrayVarTable[array])) {
					interpreter.stringArrayVarTable[array][int(parsedIndex)] = parsedNewValue
				} else {
					log.Fatalf("Interpreter error: Index '%f' out of bounds. Line %d.", parsedIndex, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.SWAP:
			array := currentToken.GetParameter(0)
			index1 := currentToken.GetParameter(1)
			index2 := currentToken.GetParameter(2)
			parsedIndex1 := -1.0
			parsedIndex2 := -1.0
			if isRawNumber, value := isRawNumber(index1); isRawNumber {
				parsedIndex1 = value
			} else if interpreter.isNumberVar(index1) {
				parsedIndex1 = interpreter.numberVarTable[index1]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", index1, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(index2); isRawNumber {
				parsedIndex2 = value
			} else if interpreter.isNumberVar(index2) {
				parsedIndex2 = interpreter.numberVarTable[index2]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", index2, currentToken.GetLine())
			}
			if interpreter.isNumberArrayVar(array) {
				if parsedIndex1 >= 0 && parsedIndex1 < float64(len(interpreter.numberArrayVarTable[array])) && parsedIndex2 >= 0 && parsedIndex2 < float64(len(interpreter.numberArrayVarTable[array])) {
					value1 := interpreter.numberArrayVarTable[array][int(parsedIndex1)]
					value2 := interpreter.numberArrayVarTable[array][int(parsedIndex2)]
					interpreter.numberArrayVarTable[array][int(parsedIndex1)] = value2
					interpreter.numberArrayVarTable[array][int(parsedIndex2)] = value1
				} else {
					log.Fatalf("Interpreter error: Index '%s' or '%s' out of bounds. Line %d.", index1, index2, currentToken.GetLine())
				}
			} else if interpreter.isStringArrayVar(array) {
				if parsedIndex1 >= 0 && parsedIndex1 < float64(len(interpreter.stringArrayVarTable[array])) && parsedIndex2 >= 0 && parsedIndex2 < float64(len(interpreter.stringArrayVarTable[array])) {
					value1 := interpreter.stringArrayVarTable[array][int(parsedIndex1)]
					value2 := interpreter.stringArrayVarTable[array][int(parsedIndex2)]
					interpreter.stringArrayVarTable[array][int(parsedIndex1)] = value2
					interpreter.stringArrayVarTable[array][int(parsedIndex2)] = value1
				} else {
					log.Fatalf("Interpreter error: Index '%s' or '%s' out of bounds. Line %d.", index1, index2, currentToken.GetLine())
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not an array variable. Line %d.", array, currentToken.GetLine())
			}
		case token.GET:
			array := currentToken.GetParameter(0)
			index := currentToken.GetParameter(1)
			parsedIndex := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(index); isRawNumber {
				parsedIndex = value
			} else if interpreter.isNumberVar(index) {
				parsedIndex = interpreter.numberVarTable[index]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s', not a number or a number variable. Line %d.", index, currentToken.GetLine())
			}
			if isRawNumberArray, value := interpreter.isRawNumberArray(array); isRawNumberArray {
				if int(parsedIndex) < len(value) {
					interpreter.numberVarTable[variableName] = value[int(parsedIndex)]
				} else {
					log.Fatalf("Interpreter error: Index out of bounds for array '%s' [%s]. Line %d.", array, index, currentToken.GetLine())
				}
				interpreter.deleteVarIfSameName(variableName, "number")
			} else if isRawStringArray, value := interpreter.isRawStringArray(array); isRawStringArray {
				if int(parsedIndex) < len(value) {
					interpreter.stringVarTable[variableName] = value[int(parsedIndex)]
				} else {
					log.Fatalf("Interpreter error: Index out of bounds for array '%s' [%s]. Line %d.", array, index, currentToken.GetLine())
				}
				interpreter.deleteVarIfSameName(variableName, "string")
			} else if interpreter.isNumberArrayVar(array) {
				if int(parsedIndex) < len(interpreter.numberArrayVarTable[array]) {
					interpreter.numberVarTable[variableName] = interpreter.numberArrayVarTable[array][int(parsedIndex)]
				} else {
					log.Fatalf("Interpreter error: Index out of bounds for array '%s' [%s]. Line %d.", array, index, currentToken.GetLine())
				}
				interpreter.deleteVarIfSameName(variableName, "number")
			} else if interpreter.isStringArrayVar(array) {
				if int(parsedIndex) < len(interpreter.stringArrayVarTable[array]) {
					interpreter.stringVarTable[variableName] = interpreter.stringArrayVarTable[array][int(parsedIndex)]
				} else {
					log.Fatalf("Interpreter error: Index out of bounds for array '%s' [%s]. Line %d.", array, index, currentToken.GetLine())
				}
				interpreter.deleteVarIfSameName(variableName, "string")
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", array, currentToken.GetLine())
			}
		case token.FOREACH:
			element := currentToken.GetParameter(0)
			array := currentToken.GetParameter(1)
			if !interpreter.foreachNamesStack.Contains(element) {
				interpreter.foreachIndexesStack.Push([]int{i, 0})
				interpreter.foreachNamesStack.Push(element)
				interpreter.lastForeachEnded = false
			}
			currentIndex := interpreter.foreachIndexesStack.Pop().([]int)[1]
			nextIndex := currentIndex + 1
			if isRawNumberArray, value := interpreter.isRawNumberArray(array); isRawNumberArray {
				if currentIndex == 0 {
					interpreter.deleteVarIfSameName(element, "number")
				}
				if currentIndex < len(value) {
					interpreter.numberVarTable[element] = value[currentIndex]
					if nextIndex < len(value) {
						interpreter.foreachIndexesStack.Push([]int{i, currentIndex + 1})
					} else {
						interpreter.lastForeachEnded = true
					}
				}
			} else if isRawStringArray, value := interpreter.isRawStringArray(array); isRawStringArray {
				if currentIndex == 0 {
					interpreter.deleteVarIfSameName(element, "string")
				}
				if currentIndex < len(value) {
					interpreter.stringVarTable[element] = value[currentIndex]
					if nextIndex < len(value) {
						interpreter.foreachIndexesStack.Push([]int{i, currentIndex + 1})
					} else {
						interpreter.lastForeachEnded = true
					}
				}
			} else if interpreter.isNumberArrayVar(array) {
				if currentIndex == 0 {
					interpreter.deleteVarIfSameName(element, "number")
				}
				if currentIndex < len(interpreter.numberArrayVarTable[array]) {
					interpreter.numberVarTable[element] = interpreter.numberArrayVarTable[array][currentIndex]
					if nextIndex < len(interpreter.numberArrayVarTable[array]) {
						interpreter.foreachIndexesStack.Push([]int{i, currentIndex + 1})
					} else {
						interpreter.varsToDelete.Push(element)
						interpreter.lastForeachEnded = true
					}
				}
			} else if interpreter.isStringArrayVar(array) {
				if currentIndex == 0 {
					interpreter.deleteVarIfSameName(element, "string")
				}
				if currentIndex < len(interpreter.stringArrayVarTable[array]) {
					interpreter.stringVarTable[element] = interpreter.stringArrayVarTable[array][currentIndex]
					if nextIndex < len(interpreter.stringArrayVarTable[array]) {
						interpreter.foreachIndexesStack.Push([]int{i, currentIndex + 1})
					} else {
						interpreter.varsToDelete.Push(element)
						interpreter.lastForeachEnded = true
					}
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", array, currentToken.GetLine())
			}
		case token.ENDFOREACH:
			goToIndex, notEmpty := interpreter.foreachIndexesStack.Top().([]int)
			if notEmpty && !interpreter.lastForeachEnded {
				i = goToIndex[0] - 1
			} else {
				interpreter.lastForeachEnded = false
				interpreter.foreachNamesStack.Pop()
				interpreter.deleteVars()
			}
		case token.RANGEARRAY:
			arrayRange := currentToken.GetParameter(0)
			parsedArrayRange := -1.0
			variableName := currentToken.GetParameter(1)
			if isRawNumber, value := isRawNumber(arrayRange); isRawNumber {
				parsedArrayRange = value
			} else if interpreter.isNumberVar(arrayRange) {
				parsedArrayRange = interpreter.numberVarTable[arrayRange]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", arrayRange, currentToken.GetLine())
			}
			interpreter.deleteVarIfSameName(variableName, "numberarray")
			array := make([]float64, int(parsedArrayRange))
			for i := 0; i < len(array); i++ {
				array[i] = float64(i)
			}
			interpreter.numberArrayVarTable[variableName] = array
		case token.RANDOM:
			min := currentToken.GetParameter(0)
			parsedMin := -1.0
			max := currentToken.GetParameter(1)
			parsedMax := -1.0
			variableName := currentToken.GetParameter(2)
			if isRawNumber, value := isRawNumber(min); isRawNumber {
				parsedMin = value
			} else if interpreter.isNumberVar(min) {
				parsedMin = interpreter.numberVarTable[min]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", min, currentToken.GetLine())
			}
			if isRawNumber, value := isRawNumber(max); isRawNumber {
				parsedMax = value
			} else if interpreter.isNumberVar(max) {
				parsedMax = interpreter.numberVarTable[max]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", max, currentToken.GetLine())
			}
			rand.Seed(time.Now().UnixNano())
			random := float64(rand.Intn(int(parsedMax-parsedMin+1))) + parsedMin
			interpreter.numberVarTable[variableName] = random
			interpreter.deleteVarIfSameName(variableName, "number")
		case token.READFILE:
			filePath := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			parsedFilePath := ""
			if isRawString, value := isRawString(filePath); isRawString {
				parsedFilePath = value
			} else if interpreter.isStringVar(filePath) {
				parsedFilePath = interpreter.stringVarTable[filePath]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", filePath, currentToken.GetLine())
			}
			parsedFilePath = filepath.Join(sourceCodeDir, parsedFilePath)
			fileContent := io.GetFileContent(parsedFilePath)
			interpreter.stringVarTable[variableName] = fileContent
			interpreter.deleteVarIfSameName(variableName, "string")
		case token.READFILELINES:
			filePath := currentToken.GetParameter(0)
			variableName := currentToken.GetParameter(1)
			parsedFilePath := ""
			if isRawString, value := isRawString(filePath); isRawString {
				parsedFilePath = value
			} else if interpreter.isStringVar(filePath) {
				parsedFilePath = interpreter.stringVarTable[filePath]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", filePath, currentToken.GetLine())
			}
			parsedFilePath = filepath.Join(sourceCodeDir, parsedFilePath)
			fileContent := io.GetFileLinesNoTrim(parsedFilePath)
			interpreter.stringArrayVarTable[variableName] = fileContent
			interpreter.deleteVarIfSameName(variableName, "stringarray")
		case token.WRITEFILE:
			filePath := currentToken.GetParameter(0)
			parsedFilePath := ""
			fileContent := currentToken.GetParameter(1)
			parsedFileContent := ""
			if isRawString, value := isRawString(filePath); isRawString {
				parsedFilePath = value
			} else if interpreter.isStringVar(filePath) {
				parsedFilePath = interpreter.stringVarTable[filePath]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", filePath, currentToken.GetLine())
			}
			if isRawString, value := isRawString(fileContent); isRawString {
				parsedFileContent = value
			} else if interpreter.isStringVar(fileContent) {
				parsedFileContent = interpreter.stringVarTable[fileContent]
			} else if isRawNumber, value := isRawNumber(fileContent); isRawNumber {
				parsedFileContent = strconv.FormatFloat(value, 'f', -1, 64)
			} else if interpreter.isNumberVar(fileContent) {
				parsedFileContent = strconv.FormatFloat(interpreter.numberVarTable[fileContent], 'f', -1, 64)
			} else if interpreter.isNumberArrayVar(fileContent) {
				for _, value := range interpreter.numberArrayVarTable[fileContent] {
					parsedFileContent += strconv.FormatFloat(value, 'f', -1, 64) + " "
				}
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", fileContent, currentToken.GetLine())
			}
			parsedFilePath = filepath.Join(sourceCodeDir, parsedFilePath)
			io.WriteToFile(parsedFilePath, parsedFileContent)
		case token.SPLIT:
			target := currentToken.GetParameter(0)
			pattern := currentToken.GetParameter(1)
			variableName := currentToken.GetParameter(2)
			parsedTarget := ""
			parsedPattern := ""
			if isRawString, value := isRawString(target); isRawString {
				parsedTarget = value
			} else if interpreter.isStringVar(target) {
				parsedTarget = interpreter.stringVarTable[target]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", target, currentToken.GetLine())
			}
			if isRawString, value := isRawString(pattern); isRawString {
				parsedPattern = value
			} else if interpreter.isStringVar(pattern) {
				parsedPattern = interpreter.stringVarTable[pattern]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", pattern, currentToken.GetLine())
			}
			splitted := strings.Split(parsedTarget, parsedPattern)
			interpreter.stringArrayVarTable[variableName] = splitted
			interpreter.deleteVarIfSameName(variableName, "stringarray")
		case token.REPLACE:
			target := currentToken.GetParameter(0)
			pattern := currentToken.GetParameter(1)
			replacement := currentToken.GetParameter(2)
			variableName := currentToken.GetParameter(3)
			parsedTarget := ""
			parsedPattern := ""
			parsedReplacement := ""
			if isRawString, value := isRawString(target); isRawString {
				parsedTarget = value
			} else if interpreter.isStringVar(target) {
				parsedTarget = interpreter.stringVarTable[target]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", target, currentToken.GetLine())
			}
			if isRawString, value := isRawString(pattern); isRawString {
				parsedPattern = value
			} else if interpreter.isStringVar(pattern) {
				parsedPattern = interpreter.stringVarTable[pattern]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", pattern, currentToken.GetLine())
			}
			if isRawString, value := isRawString(replacement); isRawString {
				parsedReplacement = value
			} else if interpreter.isStringVar(replacement) {
				parsedReplacement = interpreter.stringVarTable[replacement]
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", replacement, currentToken.GetLine())
			}
			replaced := strings.ReplaceAll(parsedTarget, parsedPattern, parsedReplacement)
			interpreter.stringVarTable[variableName] = replaced
		case token.EXIT:
			exitCode := currentToken.GetParameter(0)
			parsedExitCode := 0
			if isRawNumber, value := isRawNumber(exitCode); isRawNumber {
				parsedExitCode = int(value)
			} else if interpreter.isNumberVar(exitCode) {
				parsedExitCode = int(interpreter.numberVarTable[exitCode])
			} else {
				log.Fatalf("Interpreter error: Invalid parameter '%s'. Line %d.", exitCode, currentToken.GetLine())
			}
			os.Exit(parsedExitCode)
		}
	}
}
