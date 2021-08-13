package main

import (
	"linha-comando/app"
	"log"
	"os"
)

func main() {
	// fmt.Println("App de linha de comando para buscar IPs")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
