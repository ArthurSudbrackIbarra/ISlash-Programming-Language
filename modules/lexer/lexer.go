package lexer

import (
	"bufio"
	"islash/modules/token"
	"log"
	"os"
	"strings"
)

type Lexer struct {
	filePath string
}

func NewLexer(filePath string) *Lexer {
	return &Lexer{
		filePath: filePath,
	}
}

func (lexer *Lexer) readFileLines() []string {
	file, err := os.Open(lexer.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (lexer *Lexer) MountTokens() []*token.Token {
	tokensList := make([]*token.Token, 0)
	for _, content := range lexer.readFileLines() {
		lineFragments := strings.Split(content, " ")
		if len(lineFragments) == 0 {
			continue
		}
		switch strings.ToLower(lineFragments[0]) {
		case token.DECLARE:
			if len(lineFragments) != 3 {
				log.Fatalf("Wrong DECLARE statement, expected 2 parameters but got %d.", len(lineFragments))
			}
			tokensList = append(tokensList, token.NewToken(token.DECLARE, lineFragments[1:]))
		case token.ADDRAW:
			break
		}
	}
	return tokensList
}
