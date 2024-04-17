package main

type Subtraction struct {
	arithmicOperation
}


func (s Subtraction) Operate() Result{
	return  Result(s.Number1-s.Number2) 
}
