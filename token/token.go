package token

import "fmt"

type TokenType uint

const (
	ILLEGAL TokenType = iota
	EOF

	//literal_beg
	IDENT
	INT
	//literal_end

	//operator_beg
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
	COMMA     // ,
	SEMICOLON // ;
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	//operator_end

	//keyword_beg
	TRUE
	FALSE
	FUNC
	VAR
	IF
	ELSE
	RETURN
	//keyword_end
)

var (
	tokenTypes = map[byte]TokenType{
		'+': ADD,
		'-': SUB,
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

	tokenTypeStrings = map[TokenType]string{
		EOF:       "EOF",
		IDENT:     "IDENT",
		INT:       "INT",
		LT:        "LT",
		GT:        "GT",
		ASSIGN:    "ASSIGN",
		NOT:       "NOT",
		ADD:       "ADD",
		SUB:       "SUB",
		MUL:       "MUL",
		DIV:       "DIV",
		MOD:       "MOD",
		EQ:        "EQ",
		NEQ:       "NEQ",
		LEQ:       "LEQ",
		GEQ:       "GEQ",
		AND:       "AND",
		OR:        "OR",
		COMMA:     "COMMA",
		SEMICOLON: "SEMICOLON",
		LPAREN:    "LPAREN",
		RPAREN:    "RPAREN",
		LBRACE:    "LBRACE",
		RBRACE:    "RBRACE",
		TRUE:      "TRUE",
		FALSE:     "FALSE",
		FUNC:      "FUNC",
		VAR:       "VAR",
		IF:        "IF",
		ELSE:      "ELSE",
		RETURN:    "RETURN",
	}
)

func ToString(t TokenType) string {
	s, ok := tokenTypeStrings[t]
	if ok {
		return s
	}
	return "ILLEGAL"
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

func (this *Token) TypeIs(t TokenType) bool {
	return this.Type == t
}

func (this *Token) Eof() bool {
	return this.TypeIs(EOF)
}

func (this *Token) Illegal() bool {
	return this.TypeIs(ILLEGAL)
}

func (this *Token) String() string {
	return fmt.Sprintf("{\"type\":\"%v\",\"literal\":\"%v\"}", ToString(this.Type), this.Literal)
}

func GetTokenType(ch byte) (TokenType, bool) {
	tt, ok := tokenTypes[ch]
	return tt, ok
}
