package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	products = append(products, &product{Id: "s001", Name: "keyboard", Price: 2})
	products = append(products, &product{Id: "s002", Name: "mouse", Price: 2})
	products = append(products, &product{Id: "s003", Name: "headphone", Price: 3})
}

func main() {
	http.HandleFunc("/product", ActionProduct)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

// ActionProduct act as handler
func ActionProduct(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}

	if !AllowOnlyGET(w, r) {
		return
	}

	OutputJSON(w, GetProducts())
}

//Generate json from struct
func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
