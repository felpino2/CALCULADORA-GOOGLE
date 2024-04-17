package main

import ("log"
	"time"
	"net/http")
	
type Router struct {}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// http:// ip /path?queryString
	// queryString é parâmetro 
	// /judge?strategy=corrupt
	
	url := req.URL
	
	path := url.Path //path, o pedido - string
	
	if path != "/calculator" {
		res.Write([]byte("ERRO! 404"))
		return
	}
	
	name := url.Query().Get("Name")
	resposta := CalculatorMap[Name(name)].Operate()
	
	res.Write(resposta.Bytes())
	
}


func RunService() {
	s := &http.Server{ //struct s do tipo server 
		Addr:           "192.168.0.32:8080",
		Handler:        Router{},
		ReadTimeout:    10 * time.Second, //request
		WriteTimeout:   10 * time.Second, //response
	}
	
	log.Fatal(s.ListenAndServe())
}



