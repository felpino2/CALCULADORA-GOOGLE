package main

type Sum struct {
	arithmicOperation
}

func (Sum) Operate() Result{
	return arithmicOperation.Number1+arithmicOperation.Number2
}
