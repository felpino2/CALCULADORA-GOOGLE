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


func (s Sum) Set(n1, n2 float64){
	s.Number1 = n1
	s.Number2 = n2
}

func (s Sum) Operate() Result{ //construtor, métodoDaInterface()
	return Result(s.Number1+s.Number2) 
}


type Subtraction struct {
	ArithmicOperation
}

func (s Subtraction) Set(n1, n2 float64){
	s.Number1 = n1
	s.Number2 = n2
}	


func (s Subtraction) Operate() Result{
	return  Result(s.Number1-s.Number2) 
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
