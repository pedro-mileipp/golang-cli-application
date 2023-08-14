package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Executar vai retornar a aplicação de linha de comando pronta para ser executada
func Executar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name: "servers",
			Usage: "Busca o nome de servidores na internet",
			Flags: flags,
			Action: buscarServers,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host) // pacote que procura o IP e o Server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServers(c *cli.Context){
	host := c.String("host")
	servers, erro := net.LookupNS(host) // name server

	if erro != nil{
		log.Fatal(erro)
	}

	for _, server := range servers{
		fmt.Println(server.Host)
	}
}