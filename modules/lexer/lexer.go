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
		case token.SET:
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
		case token.POWER:
			if len(parameters) != 3 {
				log.Fatalf("Invalid POWER statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ROOT:
			if len(parameters) != 3 {
				log.Fatalf("Invalid ROOT statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
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
		case token.UPPER:
			if len(parameters) != 2 {
				log.Fatalf("Invalid UPPER statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LOWER:
			if len(parameters) != 2 {
				log.Fatalf("Invalid LOWER statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.CONTAINS:
			if len(parameters) != 3 {
				log.Fatalf("Invalid CONTAINS statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.GETCHAR:
			if len(parameters) != 3 {
				log.Fatalf("Invalid GETCHAR statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LENGTH:
			if len(parameters) != 2 {
				log.Fatalf("Invalid LENGTH statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.GREATER:
			if len(parameters) != 3 {
				log.Fatalf("Invalid GREATER statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.GREATEREQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid GREATEREQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LESS:
			if len(parameters) != 3 {
				log.Fatalf("Invalid LESS statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.LESSEQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid LESSEQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.EQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid EQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.NOTEQUAL:
			if len(parameters) != 3 {
				log.Fatalf("Invalid NOTEQUAL statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
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
		case token.RANDOM:
			if len(parameters) != 3 {
				log.Fatalf("Invalid RANDOM statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.BREAK:
			if len(parameters) != 0 {
				log.Fatalf("Invalid BREAK statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.APPEND:
			if len(parameters) != 2 {
				log.Fatalf("Invalid APPEND statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.PREPEND:
			if len(parameters) != 2 {
				log.Fatalf("Invalid PREPEND statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.REMOVEFIRST:
			if len(parameters) != 2 {
				log.Fatalf("Invalid REMOVEFIRST statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.REMOVELAST:
			if len(parameters) != 2 {
				log.Fatalf("Invalid REMOVELAST statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SETINDEX:
			if len(parameters) != 3 {
				log.Fatalf("Invalid SETINDEX statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SWAP:
			if len(parameters) != 3 {
				log.Fatalf("Invalid SWAP statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ACCESSINDEX:
			if len(parameters) != 3 {
				log.Fatalf("Invalid ACCESSINDEX statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.FOREACH:
			if len(parameters) != 2 {
				log.Fatalf("Invalid FOREACH statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.ENDFOREACH:
			if len(parameters) != 0 {
				log.Fatalf("Invalid ENDFOREACH statement, expected 0 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.RANGEARRAY:
			if len(parameters) != 2 {
				log.Fatalf("Invalid RANGEARRAY statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.READFILE:
			if len(parameters) != 2 {
				log.Fatalf("Invalid READFILE statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.READFILELINES:
			if len(parameters) != 2 {
				log.Fatalf("Invalid READFILELINES statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.WRITEFILE:
			if len(parameters) != 2 {
				log.Fatalf("Invalid WRITEFILE statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SPLIT:
			if len(parameters) != 3 {
				log.Fatalf("Invalid SPLIT statement, expected 3 parameters but got %d. Line %d.", len(parameters), line)
			}
		default:
			log.Fatalf("Invalid instruction '%s'. Line %d.", tokenType, line)
		}
		tokensList = append(tokensList, token.NewToken(line, tokenType, parameters))
	}
	return tokensList
}
