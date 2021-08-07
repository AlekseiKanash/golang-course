package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	jdata := map[string]interface{}{"headers": req.Header, "user_agent": req.UserAgent(), "uri": req.RequestURI, "host": req.Host}
	str, err := json.MarshalIndent(jdata, "", "    ")
	if err != nil {
		fmt.Printf("Error in json.Marshall: %v", err)
	}
	fmt.Fprintf(w, "%s\n", string(str))
}

func main() {
	fmt.Println("Running HTTP Listen server on 127.0.0.1:8090")
	http.HandleFunc("/", headers)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
