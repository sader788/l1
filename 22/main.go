package main

import "fmt"

type Request interface {
	GetRequest() string
}

type NormalRequest struct {
	s string
}

func (req *NormalRequest) GetRequest() string {
	return req.s
}

type ReversedRequest struct {
	s string
}

func (req *ReversedRequest) GetReversedRequest() string {
	return req.s
}

// Адаптер
type RequestAdapter struct {
	req *ReversedRequest
}

func (adapter *RequestAdapter) GetRequest() string {
	str := adapter.req.GetReversedRequest()

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Клиент
type Client struct {
}

func (c *Client) sendRequest(req Request) {
	fmt.Println(req.GetRequest())
}

func main() {
	rReq := ReversedRequest{s: "tseuqeR desreveR"}
	rAdapter := RequestAdapter{req: &rReq}

	client := Client{}
	
	client.sendRequest(&rAdapter)
}
