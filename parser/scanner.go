package parser

import (
	"fmt"
	"hermes/lexer"
	"hermes/token"
)

type scanner struct {
	toks     []*token.Token
	pos      int
	curTok   *token.Token
	peekTok  *token.Token
	peekTok2 *token.Token
	errors   []string
}

func newScanner(l *lexer.Lexer) (*scanner, error) {
	toks := l.Parse()
	if nil == toks || len(toks) < 1 {
		return nil, fmt.Errorf("no valid token")
	}
	s := &scanner{toks: toks, pos: 0, errors: []string{}}
	s.curTok = toks[0]
	sz := len(toks)
	if sz == 1 {
		s.peekTok = toks[0]
		s.peekTok2 = toks[0]
	} else if sz == 2 {
		s.peekTok = toks[1]
		s.peekTok2 = toks[1]
	} else {
		s.peekTok = toks[1]
		s.peekTok2 = toks[2]
	}
	return s, nil
}

func (this *scanner) eof() bool {
	return this.curTok.Eof()
}

func (this *scanner) peekPrecedence() int {
	return getPrecedence(this.peekTok)
}

func (this *scanner) curPrecedence() int {
	return getPrecedence(this.curTok)
}

func (this *scanner) Errors() []string {
	return this.errors
}

func (this *scanner) appendError(err string) {
	this.errors = append(this.errors, err)
}

func (this *scanner) nextToken() {
	if this.eof() {
		return
	}
	this.pos++
	this.curTok = this.toks[this.pos]
	if !this.curTok.Eof() {
		this.peekTok = this.toks[this.pos+1]
		if !this.peekTok.Eof() {
			this.peekTok2 = this.toks[this.pos+2]
		}
	}
}

func (this *scanner) expectPeek(t token.TokenType) bool {
	if this.peekTok.TypeIs(t) {
		this.nextToken()
		return true
	}
	this.appendError(fmt.Sprintf("expected next token to be %v, got %v instead", token.ToString(t), token.ToString(this.peekTok.Type)))
	return false
}

func (this *scanner) expectPeek2(t1 token.TokenType, t2 token.TokenType) bool {
	return this.peekTok.TypeIs(t1) && this.peekTok2.TypeIs(t2)
}
