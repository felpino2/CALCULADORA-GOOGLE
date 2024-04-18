package main

type Name string

var CalculatorMap = map[Name]Operating{
	"sum": Sum{},
	"+": Sum{},
	"subtraction": Subtraction{},
	"-": Subtraction{},
	
}

type Sum struct {
	ArithmicOperation
}


func (s Sum) Operate() Result{ //construtor, métodoDaInterface()
	return Result(s.Number1+s.Number2) 
}


type Subtraction struct {
	ArithmicOperation
}


func (s Subtraction) Operate() Result{
	return  Result(s.Number1-s.Number2) 
}
