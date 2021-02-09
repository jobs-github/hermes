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
	LT        // <
	GT        // >
	ASSIGN    // =
	NOT       // !
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	MOD       // %
	EQ        // ==
	NEQ       // !=
	LEQ       // <=
	GEQ       // >=
	AND       // &&
	OR        // ||
	INC       // ++
	DEC       // --
	COMMA     // ,
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	operator_end

	keyword_beg
	TRUE
	FALSE
	FUNC
	VAR
	IF
	ELSE
	RETURN
	keyword_end
)

var (
	tokenTypes = map[byte]TokenType{
		'*': MUL,
		'/': DIV,
		'%': MOD,
		',': COMMA,
		';': SEMICOLON,
		'(': LPAREN,
		')': RPAREN,
		'{': LBRACE,
		'}': RBRACE,
	}
	keywords = map[string]TokenType{
		"true":   TRUE,
		"false":  FALSE,
		"func":   FUNC,
		"var":    VAR,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}
)

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

func GetTokenType(ch byte) (TokenType, bool) {
	tt, ok := tokenTypes[ch]
	return tt, ok
}
