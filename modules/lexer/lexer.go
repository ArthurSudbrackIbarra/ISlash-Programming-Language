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
			if len(parameters) != 2 {
				log.Fatalf("Invalid ADD statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SUB:
			if len(parameters) != 2 {
				log.Fatalf("Invalid SUB statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.CONCAT:
			if len(parameters) != 2 {
				log.Fatalf("Invalid CONCAT statement, expected 2 parameters but got %d. Line %d.", len(parameters), line)
			}
		case token.SAY:
			if len(parameters) != 1 {
				log.Fatalf("Invalid SAY statement, expected 1 parameter but got %d. Line %d.", len(parameters), line)
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
