package main

import ("log"
	"time"
	"net/http"
	"encoding/json")
	
type Router struct {}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// http:// ip /path?queryString
	// queryString é parâmetro 
	// /judge?strategy=corrupt
	
	if req.Method != "GET" || path != "calculator" {
		resp := HTTPMessage{Code: 404, ShortMessage: "Not Found", Description: path + " not supported"}
		jon, _ := json.Marshal(resp)
		res.Write(jon)
		return
	}
	
	url := req.URL
	path := url.Path //path, o pedido - string
	
	operation, num1, num1, err := 
	
	op := req.URL.Query().Get("op")
	
	
	res.Write([]byte(respostaFinal))
	
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



