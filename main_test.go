package main

import "testing"

type KmExpressionTest struct {
	expression string
	variable   map[string]interface{}
	expected   bool
}

var addKmExpressionTests = []KmExpressionTest{
	{expression: "km>0", variable: map[string]interface{}{"km": 10}, expected: true},
	{expression: "km>5", variable: map[string]interface{}{"km": 4}, expected: false},
	{expression: "km<20", variable: map[string]interface{}{"km": 10}, expected: true},
	{expression: "km<5", variable: map[string]interface{}{"km": 10}, expected: false},
	{expression: "km>0 and km<20", variable: map[string]interface{}{"km": 10}, expected: true},
	{expression: "km>0 and km<5", variable: map[string]interface{}{"km": 10}, expected: false},
}

func TestKmExpressions(t *testing.T) {
	for _, test := range addKmExpressionTests {
		if output := evaluateKmExpressions(test.expression, test.variable); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}

type BaExpressionTest struct {
	expression string
	variable   map[string]interface{}
	expected   int
}

var addBaExpressionTests = []BaExpressionTest{
	{expression: "ba=ba-10", variable: map[string]interface{}{"ba": 100}, expected: 90},
	{expression: "ba=ba*0.90", variable: map[string]interface{}{"ba": 100}, expected: 90},
}

func TestBaExpressions(t *testing.T) {
	for _, test := range addBaExpressionTests {
		if output := evaluateBaExpressions(test.expression, test.variable); output != test.expected {
			t.Errorf("Output %t not equal to expected %d", output, test.expected)
		}
	}
}