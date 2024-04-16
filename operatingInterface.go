package main

type Operating interface {
	
	Operate() Result;
}

var CalculatorMap = map[string]Operating{
	"sum": Sum{},
	
}
