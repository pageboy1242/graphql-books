package main

import (
	"encoding/json"
	"fmt"
	"graphql-books/schemas"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Webserver...")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the query
	b, _ := ioutil.ReadAll(r.Body)
	query := string(b)

	res := schemas.ProcessQuery(query)

	rJSON, _ := json.Marshal(res)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}

	w.Write([]byte(rJSON))
}