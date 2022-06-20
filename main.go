package main

import (
	"net/http"

	"github.com/API_GO/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
