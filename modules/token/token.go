package token

const (
	DECLARE string = "DECLARE"
	ADD     string = "ADD"
	SUB     string = "SUB"
	SAY     string = "SAY"
)

type Token struct {
	tokenType  string
	parameters []string
}

func NewToken(tokenType string, parameters []string) *Token {
	return &Token{
		tokenType:  tokenType,
		parameters: parameters,
	}
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
