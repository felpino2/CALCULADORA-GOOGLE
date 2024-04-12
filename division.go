package main

type Division struct {
	arithmicOperation
}

func (Division) Operate() Result{
	return arithmicOperation.Number1/arithmicOperation.Number2
}
