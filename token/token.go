package token

type TokenType uint

const (
	ILLEGAL TokenType = iota
	EOF

	literal_beg
	IDENT
	INT
	literal_end

	operator_beg
	ASSIGN
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	MOD       // %
	COMMA     // ,
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	operator_end

	keyword_beg
	FUNC
	VAR
	keyword_end
)

var keywords = map[string]TokenType{
	"func": FUNC,
	"var":  VAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}
