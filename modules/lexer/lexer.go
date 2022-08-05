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
	for _, content := range fileLines {
		lineFragments := strings.Split(content, " ")
		if len(lineFragments) == 0 {
			continue
		}
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
		default:
			log.Fatalf("Invalid instruction '%s'.", tokenType)
		}
		tokensList = append(tokensList, token.NewToken(tokenType, parameters))
	}
	return tokensList
}
