package parser

import (
	"fmt"
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
	varStmt, ok := s.(*ast.VarStmt)
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
	errs := p.scanner.Errors()
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
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 3 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
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
		stmt := program.Stmts[i]
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
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 3 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	for _, stmt := range program.Stmts {
		returnStmt, ok := stmt.(*ast.ReturnStmt)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got %v", reflect.TypeOf(stmt).String())
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() != return, got %v", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentExpr(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 1 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
	}
	ident, ok := stmt.Expr.(*ast.Identifier)
	if !ok {
		t.Fatalf("Expr is not *ast.Identifier, got %v", reflect.TypeOf(stmt.Expr).String())
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value != `foobar`, got %v", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() != `foobar`, got %v", ident.TokenLiteral())
	}
}

func TestIntExpr(t *testing.T) {
	input := `5;`

	l := lexer.New(input)
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 1 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
	}
	literal, ok := stmt.Expr.(*ast.Integer)
	if !ok {
		t.Fatalf("Expr is not *ast.Identifier, got %v", reflect.TypeOf(stmt.Expr).String())
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value != 5, got %v", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral() != `5`, got %v", literal.TokenLiteral())
	}
}

func testIdentifier(t *testing.T, expr ast.Expression, value string) bool {
	ident, ok := expr.(*ast.Identifier)
	if !ok {
		t.Errorf("expr not *ast.Identifier, got %v", reflect.TypeOf(expr).String())
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value != %v, got %v", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral() != %v, got %v", value, ident.TokenLiteral())
		return false
	}
	return true
}

func testIntegerLiteral(t *testing.T, expr ast.Expression, val int64) bool {
	iv, ok := expr.(*ast.Integer)
	if !ok {
		t.Errorf("expr not *ast.IntegerLiteral, got %v", reflect.TypeOf(expr).String())
		return false
	}
	if iv.Value != val {
		t.Errorf("val not %v, got %v", val, iv.Value)
		return false
	}
	if iv.TokenLiteral() != fmt.Sprintf("%v", val) {
		t.Errorf("TokenLiteral not %v, got %v", val, iv.TokenLiteral())
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, expr ast.Expression, val bool) bool {
	iv, ok := expr.(*ast.Boolean)
	if !ok {
		t.Errorf("expr not *ast.Boolean, got %v", reflect.TypeOf(expr).String())
		return false
	}
	if iv.Value != val {
		t.Errorf("val not %v, got %v", val, iv.Value)
		return false
	}
	if iv.TokenLiteral() != fmt.Sprintf("%v", val) {
		t.Errorf("TokenLiteral not %v, got %v", val, iv.TokenLiteral())
		return false
	}
	return true
}

func TestBoolExpr(t *testing.T) {
	input := `true;`

	l := lexer.New(input)
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}

	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 1 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
	}
	literal, ok := stmt.Expr.(*ast.Boolean)
	if !ok {
		t.Fatalf("Expr is not *ast.Boolean, got %v", reflect.TypeOf(stmt.Expr).String())
	}
	if literal.Value != true {
		t.Errorf("literal.Value != true, got %v", literal.Value)
	}
	if literal.TokenLiteral() != "true" {
		t.Errorf("literal.TokenLiteral() != `true`, got %v", literal.TokenLiteral())
	}
}

func testLiteralExpression(t *testing.T, expr ast.Expression, want interface{}) bool {
	switch v := want.(type) {
	case int:
		return testIntegerLiteral(t, expr, int64(v))
	case int64:
		return testIntegerLiteral(t, expr, v)
	case string:
		return testIdentifier(t, expr, v)
	case bool:
		return testBooleanLiteral(t, expr, v)
	}
	t.Errorf("type of expr not supported: %v", reflect.TypeOf(expr).String())
	return false
}

func testInfixExpression(t *testing.T, expr ast.Expression, left interface{}, op string, right interface{}) bool {
	infixExpr, ok := expr.(*ast.InfixExpression)
	if !ok {
		t.Errorf("expr is not ast.InfixExpression, got %v", reflect.TypeOf(expr).String())
		return false
	}
	if !testLiteralExpression(t, infixExpr.Left, left) {
		return false
	}
	if infixExpr.Op != op {
		t.Errorf("infixExpr.Op != %v, got %v", op, infixExpr.Op)
		return false
	}
	if !testLiteralExpression(t, infixExpr.Right, right) {
		return false
	}
	return true
}

func TestParsingPrefixExpressions(t *testing.T) {
	cases := []struct {
		input string
		op    string
		val   interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}
	for _, tt := range cases {
		l := lexer.New(tt.input)
		p, err := New(l)
		if nil != err {
			t.Fatal(err)
		}
		program := p.ParseProgram()
		if nil == program {
			t.Fatalf("program is nil")
		}
		checkParserErrors(t, p)
		if len(program.Stmts) != 1 {
			t.Fatalf("number of program Statements: %v", len(program.Stmts))
		}

		stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
		}
		expr, ok := stmt.Expr.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expr is not *ast.PrefixExpression, got %v", reflect.TypeOf(stmt.Expr).String())
		}
		if expr.Op != tt.op {
			t.Errorf("expr.Op != %v, got %v", tt.op, expr.Op)
		}
		if testLiteralExpression(t, expr.Right, tt.val) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	cases := []struct {
		input string
		left  interface{}
		op    string
		right interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range cases {
		l := lexer.New(tt.input)
		p, err := New(l)
		if nil != err {
			t.Fatal(err)
		}
		program := p.ParseProgram()
		if nil == program {
			t.Fatalf("program is nil")
		}
		checkParserErrors(t, p)
		if len(program.Stmts) != 1 {
			t.Fatalf("number of program Statements: %v", len(program.Stmts))
		}

		stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
		}
		expr, ok := stmt.Expr.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Expr is not *ast.InfixExpression, got %v", reflect.TypeOf(stmt.Expr).String())
		}
		if !testLiteralExpression(t, expr.Left, tt.left) {
			return
		}
		if expr.Op != tt.op {
			t.Errorf("expr.Op != %v, got %v", tt.op, expr.Op)
		}
		if !testLiteralExpression(t, expr.Right, tt.right) {
			return
		}
	}
}

func TestOpPrecedParsing(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b % c", "(a + (b % c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 >= 4 == 3 <= 4", "((5 >= 4) == (3 <= 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"a && b || c", "((a && b) || c)"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
	}
	for _, tt := range cases {
		l := lexer.New(tt.input)
		p, err := New(l)
		if nil != err {
			t.Fatal(err)
		}
		program := p.ParseProgram()
		if nil == program {
			t.Fatalf("program is nil")
		}
		checkParserErrors(t, p)
		str := program.String()
		if tt.want != str {
			t.Errorf("expected %v, want %v", tt.want, str)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`
	l := lexer.New(input)
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}
	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 1 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
	}
	expr, ok := stmt.Expr.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expr is not *ast.IfExpression, got %v", reflect.TypeOf(stmt.Expr).String())
	}
	if 1 != len(expr.Clauses) {
		t.Fatalf("number of expr.Clauses: %v", len(expr.Clauses))
	}

	if !testInfixExpression(t, expr.Clauses[0].If, "x", "<", "y") {
		return
	}
	if 1 != len(expr.Clauses[0].Then.Stmts) {
		t.Fatalf("number of expr.Clauses[0].Then.Stmts: %v", len(expr.Clauses[0].Then.Stmts))
	}
	thenstmt := expr.Clauses[0].Then.Stmts[0]
	then, ok := thenstmt.(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("stmt is not *ast.ExpressionStatement, got %v", reflect.TypeOf(thenstmt).String())
	}
	if !testIdentifier(t, then.Expr, "x") {
		return
	}
	if expr.Else != nil {
		t.Fatalf("expr.Else != nil")
	}
}

func TestIfClausesExpression(t *testing.T) {
	input := `
	if (x < y) { 
		x
	} else if (x > y) {
		y
	} else {
		z
	}
	`
	l := lexer.New(input)
	p, err := New(l)
	if nil != err {
		t.Fatal(err)
	}
	program := p.ParseProgram()
	if nil == program {
		t.Fatalf("program is nil")
	}
	checkParserErrors(t, p)
	if len(program.Stmts) != 1 {
		t.Fatalf("number of program Statements: %v", len(program.Stmts))
	}

	stmt, ok := program.Stmts[0].(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got %v", reflect.TypeOf(program.Stmts[0]).String())
	}
	expr, ok := stmt.Expr.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expr is not *ast.IfExpression, got %v", reflect.TypeOf(stmt.Expr).String())
	}
	if 2 != len(expr.Clauses) {
		t.Fatalf("number of expr.Clauses: %v", len(expr.Clauses))
	}

	if !testInfixExpression(t, expr.Clauses[0].If, "x", "<", "y") {
		return
	}
	if 1 != len(expr.Clauses[0].Then.Stmts) {
		t.Fatalf("number of expr.Clauses[0].Then.Stmts: %v", len(expr.Clauses[0].Then.Stmts))
	}
	thenstmt := expr.Clauses[0].Then.Stmts[0]
	then, ok := thenstmt.(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("stmt is not *ast.ExpressionStatement, got %v", reflect.TypeOf(thenstmt).String())
	}
	if !testIdentifier(t, then.Expr, "x") {
		return
	}

	if !testInfixExpression(t, expr.Clauses[1].If, "x", ">", "y") {
		return
	}
	if 1 != len(expr.Clauses[1].Then.Stmts) {
		t.Fatalf("number of expr.Clauses[0].Then.Stmts: %v", len(expr.Clauses[1].Then.Stmts))
	}
	thenstmt2 := expr.Clauses[1].Then.Stmts[0]
	then2, ok := thenstmt2.(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("stmt is not *ast.ExpressionStatement, got %v", reflect.TypeOf(thenstmt2).String())
	}
	if !testIdentifier(t, then2.Expr, "y") {
		return
	}

	if expr.Else == nil {
		t.Fatalf("expr.Else == nil")
	}

	if 1 != len(expr.Else.Stmts) {
		t.Fatalf("number of expr.Else.Stmts: %v", len(expr.Else.Stmts))
	}
	elsestmt := expr.Else.Stmts[0]
	thenelse, ok := elsestmt.(*ast.ExpressionStmt)
	if !ok {
		t.Fatalf("elsestmt is not *ast.ExpressionStatement, got %v", reflect.TypeOf(elsestmt).String())
	}
	if !testIdentifier(t, thenelse.Expr, "z") {
		return
	}
}
