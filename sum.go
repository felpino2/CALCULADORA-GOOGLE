package main

type Sum struct {
	arithmicOperation
}

func (s Sum) Constructor(numero1 float32, numero2 float32) {

	s.Number1 = numero1
	s.Number2 = numero2
	return s
}

func (Sum) Operate(float32, float32) Result{
	return arithmicOperation.Number1+arithmicOperation.Number2
}
