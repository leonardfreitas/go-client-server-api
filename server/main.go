package main

import (
	"net/http"

	"github.com/leonardfreitas/go-client-server-api/server/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", controllers.GetDollarControler)
	http.ListenAndServe(":8080", mux)
}
