package main

import "strconv"

type Result float64

func (r Result) Bytes() []byte { //TRANSFORMA EM BYTES

	fmt64 := strconv.FormatFloat(float64(r), 'f', -1, 64)
	return []byte(fmt64)
}
