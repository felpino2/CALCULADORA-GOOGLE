package main

type Subtraction struct {
	arithmicOperation
}

func (Subtraction) Operate() Result{
	return arithmicOperation.Number1-arithmicOperation.Number2
}
