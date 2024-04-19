//http://192.168.0.32:8080/calculator?op=sum&n1=2&n2=5

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
	
	
	url := req.URL
	path := url.Path
	
	if req.Method != "GET" || path != "/calculator" {
		resp := HTTPMessageMap[404]
		jon, _ := json.Marshal(resp)
		res.Write(jon)
		return
	}
	
	
	operation, numb1, numb2, erro := Translator(req)
	
	if erro != nil {
		resp := HTTPMessageMap[500]
		jon, _ := json.Marshal(resp)
		res.Write(jon)
		return
	}
	
	result, err := Operation(operation, numb1, numb2, res)
	
	if err != nil {
		//codigo de erro
		return
	}
	
	resp := SuccessfulOp200 (result, operation)
	jon, _ := json.Marshal(resp)
	res.Write(jon)
		
	
	
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



