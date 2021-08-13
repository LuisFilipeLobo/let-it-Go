package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Vai retornar o app pronto para execução
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores"

	// delcaração das flags
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	// slice de comandos para o app executar
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar IPs",
			Flags:  flags,
			Action: buscarIp,
		},
		{
			Name:   "server",
			Usage:  "Buscar nomes de Servidores",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return app
}

// pegar host passado como flag e buscar IPs com LookupIP()
func buscarIp(ctx *cli.Context) {
	host := ctx.String("host")

	ips, erro := net.LookupIP(host)

	// encerrar app em caso de erro
	if erro != nil {
		log.Fatal(erro)
	}

	// imprimir todos os IPs
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// pegar host passado como flag e buscar nome(s) de servidor(es) com LookupNS()
func buscarServer(srv *cli.Context) {
	host := srv.String("host")

	server, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	// imprimir
	for _, srv := range server {
		fmt.Println(srv.Host)
	}
}
