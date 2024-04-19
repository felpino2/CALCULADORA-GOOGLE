package main

import "math"

type Name string

var CalculatorMap = map[Name]Operating{
	"sum": &Sum{},
	"+": &Sum{},
	"sub": &Subtraction{},
	"-": &Subtraction{},
	"/": &Division{},
	"div": &Division{},
	"*": &Multiply{},
	"mul": &Multiply{},
	"root": &Root{},
	"&": &Root{}, 
	
}

func (a *ArithmicOperation) Set(n1, n2 float64) {
	a.Number1 = n1
	a.Number2 = n2
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


type Division struct {
	ArithmicOperation
}

func (s Division) Operate() Result {
	return Result(s.Number1/s.Number2)
}

type Multiply struct {
	ArithmicOperation
}
func (s Multiply) Operate() Result{
	return Result(s.Number1*s.Number2)
}

type Root struct {
	ArithmicOperation
}

func (s Root) Operate() Result{
	return Result(math.Pow(s.Number1, (1/s.Number2)))
} 


func Operation(operation string, n1, n2 float64) (Result, error){
	valor, ok := CalculatorMap[Name(operation)]
	
	if !ok {
	//ADICIONAR FUNÇÃO PARA CODIGO 404
		return Result(0), nil
	}
	
	valor.Set(n1, n2)
	
	resultado := valor.Operate()
	
	return resultado, nil	
}
