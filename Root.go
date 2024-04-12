package main

import "math"

type Root struct {
	arithmicOperation
}

func (Root) Operate() Result{
	return math.Pow(arithmicOperation.Number1, (1/arithmicOperation.Number2))
}
