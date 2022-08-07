package token

const (
	DECLARE      string = "DECLARE"
	ADD          string = "ADD"
	SUB          string = "SUB"
	MULT         string = "MULT"
	DIV          string = "DIV"
	MOD          string = "MOD"
	GREATER      string = "GREATER"
	GREATEREQUAL string = "GREATEREQUAL"
	LESSER       string = "LESSER"
	LESSEREQUAL  string = "LESSEREQUAL"
	EQUAL        string = "EQUAL"
	CONCAT       string = "CONCAT"
	SAY          string = "SAY"
	WHILE        string = "WHILE"
	ENDWHILE     string = "ENDWHILE"
)

type Token struct {
	line       int
	tokenType  string
	parameters []string
}

func NewToken(line int, tokenType string, parameters []string) *Token {
	return &Token{
		line:       line,
		tokenType:  tokenType,
		parameters: parameters,
	}
}

func (token *Token) GetLine() int {
	return token.line
}

func (token *Token) GetType() string {
	return token.tokenType
}

func (token *Token) GetParameter(index int) string {
	if index < 0 || index >= len(token.parameters) {
		return ""
	}
	return token.parameters[index]
}
