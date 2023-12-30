package main

import "net/http"

func main() {
	InitLoad(LoadRequest{
		ReqTotal: 200,
		Strategy: Linear,
		Time:     "60s",
		Request: Request{
			URL:    "https://vaibhtest.requestcatcher.com/test",
			Method: http.MethodGet,
			Body:   nil,
		},
	})
}
