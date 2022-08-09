package token

const (
	DECLARE          string = "DECLARE"
	ADD              string = "ADD"
	SUB              string = "SUB"
	MULT             string = "MULT"
	DIV              string = "DIV"
	MOD              string = "MOD"
	INCREMENT        string = "INCREMENT"
	DECREMENT        string = "DECREMENT"
	GREATERTHAN      string = "GREATERTHAN"
	GREATERTHANEQUAL string = "GREATERTHANEQUAL"
	LESSTHAN         string = "LESSTHAN"
	LESSTHANEQUAL    string = "LESSTHANEQUAL"
	NOT              string = "NOT"
	AND              string = "AND"
	OR               string = "OR"
	IF               string = "IF"
	ELSE             string = "ELSE"
	ENDIF            string = "ENDIF"
	EQUAL            string = "EQUAL"
	NOTEQUAL         string = "NOTEQUAL"
	CONCAT           string = "CONCAT"
	GETCHAR          string = "GETCHAR"
	LENGTH           string = "LENGTH"
	SAY              string = "SAY"
	INPUT            string = "INPUT"
	WHILE            string = "WHILE"
	ENDWHILE         string = "ENDWHILE"
	BREAK            string = "BREAK" // TBD...
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
