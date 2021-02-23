package parser

import (
	"hermes/ast"
	"hermes/lexer"
	"reflect"
	"testing"
)

func testVarStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral() not `var`, got=%v", s.TokenLiteral())
		return false
	}
	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s is not *ast.VarStatement, got=%v", reflect.TypeOf(s).String())
		return false
	}
	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value != %v, got=%v", name, varStmt.Name.Value)
		return false
	}
	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name.TokenLiteral() != %v, got=%v", name, varStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()
	if len(errs) < 1 {
		return
	}
	t.Errorf("parser has %v errors", len(errs))
	for _, msg := range errs {
		t.Errorf("parser error: %v", msg)
	}
	t.FailNow()
}

func TestVarStatements(t *testing.T) {
	input := `
	var x = 5;
	var y = 10;
	var foobar = 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("number of program Statements: %v", len(program.Statements))
	}

	tests := []struct {
		name string
		want string
	}{
		{"1", "x"},
		{"2", "y"},
		{"3", "foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatements(t, stmt, tt.want) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("number of program Statements: %v", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got %v", reflect.TypeOf(stmt).String())
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() != return, got %v", returnStmt.TokenLiteral())
		}
	}
}
