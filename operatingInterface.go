package main

type Operating interface {
	
	Operate() Result; //método da interface
}

type Name string

var CalculatorMap = map[Name]Operating{
	"sum": Sum{},
	"subtraction": Subtraction{},
	
}
