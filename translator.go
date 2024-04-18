package main

import (
	"strings"
)

func Translator(string(fullQuery) string) string{
	
	var queryArray []string
	
	for k := range fullQuery{
	     if _, ok := CalculatorMap[fullQuery[k]]; ok {
	     	queryArray = strings.Split(fullQuery, fullQuery[k])
	     }
	}
	
	
}	
