package main

type arithmicOperation struct {
	Number1 float32 
	Number2 float32
}

CalculatorMap = map[string]Operating{
	"sum": Sum{},
	"subtraction": Subtraction{},
	"multiply": Multiply{},
	"division": Division{},
	"pow": Pow{}
	"root": Root{}
}



