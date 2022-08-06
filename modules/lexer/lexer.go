package lexer

import (
	"islash/modules/io"
	"islash/modules/token"
	"log"
	"strings"
)

func MountTokens(filePath string) []*token.Token {
	tokensList := make([]*token.Token, 0)
	fileLines := io.GetFileLines(filePath)
	for index, content := range fileLines {
		if content == "" {
			continue
		}
		quoted := false
		lineFragments := strings.FieldsFunc(content, func(char rune) bool {
			if char == '"' {
				quoted = !quoted
			}
			return !quoted && char == ' '
		})
		line := index + 1
		tokenType := strings.ToUpper(lineFragments[0])
		parameters := lineFragments[1:]
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
		default:
			log.Fatalf("Invalid instruction '%s'.", tokenType)
		}
		tokensList = append(tokensList, token.NewToken(line, tokenType, parameters))
	}
	return tokensList
}
