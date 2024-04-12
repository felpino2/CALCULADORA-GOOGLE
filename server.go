package main

import ("log"
	"time"
	"net/http")
	
type Router struct {}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL
	
	path := url.Path
	
	if path != "/calculator" {
		
	}
}


func RunService() {
	s := &http.Server{
		Addr:           "172.22.51.30:8080",
		Handler:        Router{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	
	log.Fatal(s.ListenAndServe())
}



