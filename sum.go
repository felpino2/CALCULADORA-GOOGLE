package main

type Sum struct {
	arithmicOperation
}


func (s Sum) Operate() Result{
	return Result(s.Number1+s.Number2)
}
