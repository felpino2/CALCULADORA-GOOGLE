package main

import ("log"
	"time"
	"net/http"
	"encoding/json")
	
type Router struct {}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
		/* http:// ip /path?queryString
	 		queryString é parâmetro 
	 		judge?strategy=corrupt */
	
	
	
	
	if req.Method != "GET"  {
		resp := HTTPMessage{Code: 404, ShortMessage: "Not Found"}
		jon, _ := json.Marshal(resp)
		res.Write(jon)
		return
	}
	
	
	operation, numb1, numb2, erro := Translator(req)
	
	if erro != nil {
		//codigo de rro
		return
	}
	
	result, err := Operation(operation, numb1, numb2)
	
	if err != nil {
		//codigo de erro
		return
	}
	
	res.Write(result.Bytes())
	
	
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



