package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonData struct {
	Host      string
	UserAgent string
	Url       string
	Headers   http.Header
}

func headers(w http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("%s%s", req.Host, req.URL.Path)
	jdata := JsonData{Host: req.Host, Headers: req.Header, UserAgent: req.UserAgent(), Url: url}
	str, err := json.MarshalIndent(jdata, "", "    ")
	if err != nil {
		fmt.Printf("Error in json.Marshall: %v", err)
	}
	fmt.Fprintf(w, "%s\n", string(str))
}

func main() {
	http.HandleFunc("/", headers)

	http.ListenAndServe(":8090", nil)
}
