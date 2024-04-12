package main

import "math"

type Pow struct {
	arithmicOperation
}

func (Pow) Operate() Result{
	return math.Pow(arithmicOperation.Number1, arithmicOperation.Number2)
}
