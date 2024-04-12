package main

type Multiply struct {
	arithmicOperation
}

func (Multiply) Operate() Result{
	return arithmicOperation.Number1*arithmicOperation.Number2
}
