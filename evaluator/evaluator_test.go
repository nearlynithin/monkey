package evaluator

import (
	"testing"

	"github.com/nearlynithin/monkey/lexer"
	"github.com/nearlynithin/monkey/object"
	"github.com/nearlynithin/monkey/parser"
)


func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input string
		expected int64
	}{
		{"5",5},
		{"8",8},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t,evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj is not *object.Integer, got=%T",obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has a wrong value, expected=%d, got=%d",expected,result.Value)
		return false
	}

	return true
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{"true",true},
		{"false",false},
	}

	for _,tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t,evaluated,tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj  object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not *object.Boolean, got=%T",obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("result.Value is wrong, expected=%t, but got=%t",expected,result.Value)
		return false
	}

	return true

}