package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Return  aplication command line
func Gerar() *cli.App {

	app := cli.NewApp()

	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar Ips e Nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar ip na net",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar o nome dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// Ota função
func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Ota função 2

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
