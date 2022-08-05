package token

/*
	Token identifiers.
*/
const (
	DECLARE string = "declare"
	ADDRAW         = "addraw"
	ADDVAR         = "addvar"
	SUBRAW         = "subraw"
	SUBVAR         = "subvar"
)

type Token struct {
	identifier string
	parameters []string
}

func NewToken(identifier string, parameters []string) *Token {
	return &Token{
		identifier: identifier,
		parameters: parameters,
	}
}
