package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(handler http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	fmt.Printf("%+v\n", req.Header.Get("X-Truffle-Signature"))
	fmt.Printf("%+v\n", string(data))
}
