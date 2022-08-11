package token

const (
	SET          string = "SET"
	ADD          string = "ADD"
	SUB          string = "SUB"
	MULT         string = "MULT"
	DIV          string = "DIV"
	MOD          string = "MOD"
	POWER        string = "POWER"
	ROOT         string = "ROOT"
	INCREMENT    string = "INCREMENT"
	DECREMENT    string = "DECREMENT"
	GREATER      string = "GREATER"
	GREATEREQUAL string = "GREATEREQUAL"
	LESS         string = "LESS"
	LESSEQUAL    string = "LESSEQUAL"
	NOT          string = "NOT"
	AND          string = "AND"
	OR           string = "OR"
	IF           string = "IF"
	ELSE         string = "ELSE"
	ENDIF        string = "ENDIF"
	EQUAL        string = "EQUAL"
	NOTEQUAL     string = "NOTEQUAL"
	CONCAT       string = "CONCAT"
	GETCHAR      string = "GETCHAR"
	LENGTH       string = "LENGTH"
	SAY          string = "SAY"
	INPUT        string = "INPUT"
	WHILE        string = "WHILE"
	ENDWHILE     string = "ENDWHILE"
	BREAK        string = "BREAK"
	FOREACH      string = "FOREACH"
	ENDFOREACH   string = "ENDFOREACH"
	ACCESSINDEX  string = "ACCESSINDEX"
	APPEND       string = "APPEND"
	PREPEND      string = "PREPEND"
	REMOVEFIRST  string = "REMOVEFIRST"
	REMOVELAST   string = "REMOVELAST"
	SWAP         string = "SWAP"
	SETINDEX     string = "SETINDEX"
	RANGEARRAY   string = "RANGEARRAY"
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
