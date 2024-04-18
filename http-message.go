package main


type HTTPMessage struct {
	Code int `json:"code"`
	ShortMessage string `json:"smsg"`
	Description string `json:"desc"`
}
