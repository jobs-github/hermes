package evaluator

import (
	"hermes/lexer"
	"hermes/object"
	"hermes/parser"
	"testing"
)

func testEval(input string) (object.Object, error) {
	l := lexer.New(input)
	p, err := parser.New(l)
	if nil != err {
		return nil, err
	}
	program := p.ParseProgram()
	return program.Eval(), nil
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

func TestEvalIntegerExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOp(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated, err := testEval(tt.input)
		if nil != err {
			t.Fatal(err)
		}
		testBooleanObject(t, evaluated, tt.expected)
	}
}
