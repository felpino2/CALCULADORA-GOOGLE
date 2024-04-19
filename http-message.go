package main


type HTTPMessage struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

type OperationMessage struct {
	Code int `json:"code"`
	Result string `json:"res"`
	Operation string `json:"op"`
}


var HTTPMessageMap = map[int]HTTPMessage {
	500 : {Code: 500, Error: "Something went wrong"},
	404 : {Code: 404, Error: "Not found"},
}

func InvalidOp400 (operation string) OperationMessage {
	return OperationMessage{Code: 400, Result: "Invalid Operation", Operation: operation}
}

/*func SuccessfulOp200 () OperationMessage {
	return 
} */
