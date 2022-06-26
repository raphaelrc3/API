package main

import (
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Inicializando a API")

	r := router.Gerar()

	log.fatal(http.ListenAndServe(":8000, r"))

}
