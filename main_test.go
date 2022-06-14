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
	for i, test := range addKmExpressionTests {
		if output := evaluateBooleanExpressions(test.expression, test.variable); output != test.expected {
			t.Errorf("%d Output %t not equal to expected %t", i, output, test.expected)
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
	{expression: "BA=BA-10", variable: map[string]interface{}{"BA": 100}, expected: 90}, //case sensitive
	{expression: "BA=BA-1", variable: map[string]interface{}{"BA": 1000000000000}, expected: 999999999999}, //with big values
	{expression: "ba=ba*0.90", variable: map[string]interface{}{"ba": 100}, expected: 90},
	{expression: "ba=ba*0.90", variable: map[string]interface{}{"ba": 100000000000000}, expected: 90000000000000},
}

func TestBaExpressions(t *testing.T) {
	for i, test := range addBaExpressionTests {
		if output := evaluateExpressionWithNumber(test.expression, test.variable); output != test.expected {
			t.Errorf("%d Output %d not equal to expected %d", i, output, test.expected)
		}
	}
}