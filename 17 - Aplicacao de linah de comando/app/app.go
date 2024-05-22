package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a plicao de linha de comando para ser executada
func Geral() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereco na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca servidores na intenrnet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
