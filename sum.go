package main

type Sum struct {
	arithmicOperation
}


func (s Sum) Operate() Result{ //construtor, métodoDaInterface()
	return Result(s.Number1+s.Number2) 
}
