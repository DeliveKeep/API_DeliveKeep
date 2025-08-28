package main

import (
	"API/src/config"
	"API/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	r := routes.Rotear()

	fmt.Printf("Escutando na porta %d", config.PortaAPI)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortaAPI), r))
}
