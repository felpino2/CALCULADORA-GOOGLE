package main

type Operating interface {
	
	Operate(float32, float32) Result
}

var CalculatorMap = map[string]Operating{
	"sum": Sum{},
	"subtraction": Subtraction{},
	"multiply": Multiply{},
	"division": Division{},
	"pow": Pow{}
	"root": Root{}
}
