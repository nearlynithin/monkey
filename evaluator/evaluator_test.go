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
		{"-8",-8},
		{"-10",-10},
		{"5 + 5 + 5 + 5 -10",10},
		{"2 * 2 * 2 * 2 * 2",32},
		{"-50 + 100 + -50",0},
		{"5 * 2 + 10",20},
		{"5 + 2 * 10",25},
		{"20 + 2 * -10",0},
		{"50 / 2 * 2 + 10",60},
		{"2 * (5 + 10)",30},
		{"3 * 3 * 3 + 10",37},
		{"( 5 + 10 * 2 + 15 / 3) * 2 + -10",50},
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
		{"1 < 2",true},
		{"1 > 2",false},
		{"2 == 2",true},
		{"2 == 3",false},
		{"2 != 3",true},
		{"2 != 2",false},
		{"true == true",true},
		{"false == false",true},
		{"true != false",true},
		{"false != true",true},
		{"true == false",false},
		{"(1 > 2) == true",false},
		{"(1 > 2) == false",true},
		{"(1 < 2) == true",true},
		{"(1 < 2) == false",false},
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

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{"!true",false},
		{"!false",true},
		{"!5",false},
		{"!!true",true},
		{"!!false",false},
		{"!!5",true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t,evaluated,tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct{
		input string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1<2) { 10 }", 10},
		{"if (1>2) { 10 }", nil},
		{"if (1<2) { 10 } else { 5 }", 10},
		{"if (1>2) { 10 } else { 5 }", 5},
	}

	for _,tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t,evaluated,int64(integer))
		} else {
			testNullObject(t,evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("evaluated object is not NULL, got=%T (%+v)",obj,obj)
		return false
	}
	return true
}