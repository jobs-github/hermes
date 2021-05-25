package main

import (
	"Q/lexer"
	"Q/object"
	"Q/parser"
	"reflect"
	"testing"
)

func testEval(input string) (object.Object, error) {
	l := lexer.New(input)
	p, err := parser.New(l)
	if nil != err {
		return nil, err
	}
	program := p.ParseProgram()
	env := object.NewEnv()
	return program.Eval(env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not integer, got=%v", obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, got=%v, want: %v", result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	_, ok := obj.(*object.Null)
	if !ok {
		t.Errorf("object is not null, got=%v", obj)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not boolean, got=%v", obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, got=%v, want: %v", result.Value, expected)
		return false
	}
	return true
}

func testEvalObject(t *testing.T, evaluated object.Object, expected interface{}) {
	switch et := expected.(type) {
	case bool:
		testBooleanObject(t, evaluated, et)
	case object.Null:
		testNullObject(t, evaluated)
	case int:
		testIntegerObject(t, evaluated, int64(et))
	}
}

func TestEvalExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!null", true},
		{"!!null", false},

		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"15 % 10", 5},

		{"1 && 2", 2},
		{"2 && 1", 1},
		{"0 && 2", 0},
		{"1 || 2", 1},
		{"2 || 1", 2},
		{"0 || 2", 2},

		{"true + 1", 2},
		{"false + 1", 1},
		{"true - 1", 0},
		{"false - 1", -1},
		{"true * 2", 2},
		{"false * 2", 0},
		{"true / 2", 0},
		{"false / 2", 0},
		{"true % 2", 1},
		{"false % 2", 0},

		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 >= 1", true},
		{"1 <= 1", true},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},

		{"true > 1", false},
		{"true >= 1", true},
		{"true < 1", false},
		{"true <= 1", true},
		{"true == 1", true},
		{"false > 0", false},
		{"false >= 0", true},
		{"false < 0", false},
		{"false <= 0", true},
		{"false == 0", true},

		{"true && 2", 2},
		{"2 && true", 1},
		{"false && 2", 0},
		{"true || 2", 1},
		{"2 || true", 2},
		{"false || 2", 2},

		{"null > 0", false},
		{"null >= 0", false},
		{"null < 0", true},
		{"null <= 0", true},
		{"null != 0", true},
		{"null == 0", false},
		{"null && 0", object.Nil},
		{"null || 0", 0},

		{"null > false", false},
		{"null >= false", false},
		{"null < false", true},
		{"null <= false", true},
		{"null != false", true},
		{"null == false", false},
		{"null && false", object.Nil},
		{"null || false", false},

		{"0 > null", true},
		{"0 >= null", true},
		{"0 < null", false},
		{"0 <= null", false},
		{"0 != null", true},
		{"0 == null", false},
		{"0 && null", 0},
		{"1 && null", object.Nil},
		{"0 || null", object.Nil},
		{"1 || null", 1},

		{"false > null", true},
		{"false >= null", true},
		{"false < null", false},
		{"false <= null", false},
		{"false != null", true},
		{"false == null", false},
		{"false && null", false},
		{"true && null", object.Nil},
		{"false || null", object.Nil},
		{"true || null", true},

		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"true && true", true},
		{"true && false", false},
		{"false && false", false},
		{"true || true", true},
		{"true || false", true},
		{"false || false", false},
		{"true + true", 2},
		{"true - true", 0},
		{"true * true", 1},
		{"true / true", 1},
		{"true % true", 0},
		{"true > true", false},
		{"true >= true", true},
		{"true < true", false},
		{"true <= true", true},

		{"null == null", true},
		{"null != null", false},
		{"null > null", false},
		{"null >= null", true},
		{"null < null", false},
		{"null <= null", true},
		{"null && null", object.Nil},
		{"null || null", object.Nil},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testEvalObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (2 < 1) { 10 } else if (2 < 2) { 20 } else { 30 }", 30},
		{"if (2 < 1) { 10 } else if (2 > 1) { 20 } else { 30 }", 20},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testEvalObject(t, evaluated, tt.expected)
	}
}

func TestReturnStmts(t *testing.T) {
	stmt := `
	if (10 > 1) {
		if (10 > 1) {
			return 10;
		}
		return 1;
	}
	`
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9", 10},
		{"9; return 2 * 5; 9", 10},
		{stmt, 10},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testEvalObject(t, evaluated, tt.expected)
	}
}

func TestVarStmts(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"var a = 5; a;", 5},
		{"var a = 5 * 5; a;", 25},
		{"var a = 5; var b = a; b;", 5},
		{"var a = 5; var b = a; var c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testEvalObject(t, evaluated, tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "func(x) { x + 2; };"
	evaluated, err := testEval(input)
	if nil != err {
		t.Fatal(err)
	}
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not function, got %v", reflect.TypeOf(evaluated).String())
	}
	arguments := len(fn.Args)
	if arguments != 1 {
		t.Fatalf("function has wrong args, got %v", arguments)
	}
	argument := fn.Fn.ArgumentOf(0)
	if argument != "x" {
		t.Fatalf("argument of 0 not x, got `%v`", argument)
	}
	body := fn.Fn.Body()
	expected := "(x + 2)"
	if body != expected {
		t.Fatalf("body not (x + 2), got `%v`", body)
	}
}

func TestFunctionCases(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"var identity = func(x) { x; }; identity(5)", 5},
		{"var identity = func(x) { return x; }; identity(5)", 5},
		{"var double = func(x) { x * 2; }; double(5)", 10},
		{"var add = func(x, y) { x + y; }; add(5, 5)", 10},
		{"var add = func(x, y) { x + y; }; add(5 + 5, add(5, 5))", 20},
		{"func(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testEvalObject(t, evaluated, tt.expected)
	}
}
