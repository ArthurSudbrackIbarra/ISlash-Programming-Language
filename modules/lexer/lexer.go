package lexer

import (
	"islash/modules/io"
	"islash/modules/token"
	"log"
	"strings"
)

func splitInstruction(lineContent string) []string {
	quote := false
	splittedInstruction := strings.FieldsFunc(lineContent, func(char rune) bool {
		if char == '"' {
			quote = !quote
		}
		return !quote && char == ' '
	})
	return splittedInstruction
}

func MountTokens(filePath string) []*token.Token {
	tokensList := make([]*token.Token, 0)
	fileLines := io.GetFileLines(filePath)
	for index, lineContent := range fileLines {
		if lineContent == "" || strings.HasPrefix(lineContent, "#") {
			continue
		}
		splittedInstruction := splitInstruction(lineContent)
		line := index + 1
		tokenType := strings.ToUpper(splittedInstruction[0])
		parameters := make([]string, 0)
		if len(splittedInstruction) > 1 {
			parameters = splittedInstruction[1:]
		}
		switch tokenType {
		case token.DECLARE:
			if len(parameters) != 2 {
				log.Fatalf("Invalid DECLARE statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ADD:
			if len(parameters) != 3 {
				log.Fatalf("Invalid ADD statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SUB:
			if len(parameters) != 3 {
				log.Fatalf("Invalid SUB statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.MULT:
			if len(parameters) != 3 {
				log.Fatalf("Invalid MULT statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.DIV:
			if len(parameters) != 3 {
				log.Fatalf("Invalid DIV statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.MOD:
			if len(parameters) != 3 {
				log.Fatalf("Invalid MOD statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.INCREMENT:
			if len(parameters) != 1 {
				log.Fatalf("Invalid INCREMENT statement, expected 1 parameter but got %d. Line %d.", len(parameters), line)
			}
		case token.DECREMENT:
			if len(parameters) != 1 {
				log.Fatalf("Invalid DECREMENT statement, expected 1 parameter but got %d. Line %d.", len(parameters), line)
			}
		case token.NOT:
			if len(parameters) != 2 {
				log.Fatalf("Invalid NOT statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.AND:
			if len(parameters) != 3 {
				log.Fatalf("Invalid AND statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.OR:
			if len(parameters) != 3 {
				log.Fatalf("Invalid OR statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.CONCAT:
			if len(parameters) != 2 {
				log.Fatalf("Invalid CONCAT statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.GREATERTHAN:
			if len(parameters) != 3 {
				log.Fatalf("Invalid GREATER statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.GREATERTHANEQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid GREATEREQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LESSTHAN:
			if len(parameters) != 3 {
				log.Fatalf("Invalid LESSER statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LESSTHANEQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid LESSEREQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.EQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid EQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.IF:
			if len(parameters) != 1 {
				log.Fatalf("Invalid IF statement, expected 1 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ELSE:
			if len(parameters) != 0 {
				log.Fatalf("Invalid ELSE statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ENDIF:
			if len(parameters) != 0 {
				log.Fatalf("Invalid ENDIF statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SAY:
			if len(parameters) != 1 {
				log.Fatalf("Invalid SAY statement, expected 1 parameter but got %d. Line %d.", len(parameters), line)
			}
		case token.INPUT:
			if len(parameters) != 2 {
				log.Fatalf("Invalid INPUT statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.WHILE:
			if len(parameters) != 1 {
				log.Fatalf("Invalid WHILE statement, expected 1 parameter but got %d. Line %d.", len(parameters), line)
			}
		case token.ENDWHILE:
			if len(parameters) != 0 {
				log.Fatalf("Invalid ENDWHILE statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		default:
			log.Fatalf("Invalid instruction '%s'.", tokenType)
		}
		tokensList = append(tokensList, token.NewToken(line, tokenType, parameters))
	}
	return tokensList
}
