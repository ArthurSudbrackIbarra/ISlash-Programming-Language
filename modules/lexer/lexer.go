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
		quoted := false
		lineFragments := strings.FieldsFunc(content, func(r rune) bool {
			if r == '"' {
				quoted = !quoted
			}
			return !quoted && r == ' '
		})
		tokenType := strings.ToUpper(lineFragments[0])
		parameters := lineFragments[1:]
		switch tokenType {
		case token.DECLARE:
			if len(parameters) != 2 {
				log.Fatalf("Invalid DECLARE statement, expected 2 parameters but got %d.", len(parameters))
			}
		case token.ADD:
			if len(parameters) != 2 {
				log.Fatalf("Invalid ADD statement, expected 2 parameters but got %d.", len(parameters))
			}
		case token.SUB:
			if len(parameters) != 2 {
				log.Fatalf("Invalid SUB statement, expected 2 parameters but got %d.", len(parameters))
			}
		case token.SAY:
			if len(parameters) != 1 {
				log.Fatalf("Invalid SAY statement, expected 1 parameter but got %d.", len(parameters))
			}
		default:
			log.Fatalf("Invalid instruction '%s'.", tokenType)
		}
		tokensList = append(tokensList, token.NewToken(index+1, tokenType, parameters))
	}
	return tokensList
}
