package main

import (
	"net/http"
	"strconv"
)

func Translator(req *http.Request) (string, float64, float64, error){

	tradutorURL := req.URL.Query()
	
	
	operation := tradutorURL.Get("op")
	numero1 := tradutorURL.Get("n1")
	numero2 := tradutorURL.Get("n2")
	
	if operation == "" || numero1 == "" || numero2 == "" {
		return "", 0, 0, nil
	}
	
	floatNumero1, err1 := strconv.ParseFloat(numero1, 64)
	floatNumero2, err2 := strconv.ParseFloat(numero2, 64)
	
	if err1 != nil || err2 != nil {
		return "", 0, 0, nil 
	}
	
	return operation, floatNumero1, floatNumero2, nil
	
}	
