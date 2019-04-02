package main

import (
  "log"
  "os"
  "fmt"
  "./request"
  "github.com/urfave/cli"
  "strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_net"
	app.Usage = "cli_net [global options] command [command options] [arguments...]"

	var pr request.Params
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			Value:       "server",
			Destination: &pr.Server,
		},
		cli.StringFlag{
			Name:        "port",
			Value:       "8080",
			Destination: &pr.Port,
		},
	}

	getLatestServer := func() request.Server {
		srv := request.NewServer(pr)
		err := srv.UpdateVersion()
		if err != nil {
			log.Fatal(err)
		}
		return srv
	}

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "Returns all available interfaces on the host and virtual network",
			Action: func(c *cli.Context) {
				srv := getLatestServer()
				listOfNames, err := request.InterfaceListRequest(srv)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println(strings.Join(listOfNames.AllIntr, ","))
			},
		},
		{
			Name:  "show",
			Usage: "Show info about interface with <name>",
			Action: func(c *cli.Context) {
				if c.NArg() == 0 {
					fmt.Println("not enough args")
					return
				}

				srv := getLatestServer()
				name := c.Args()[0]
				netIf, err := request.InterfaceInfoRequest(name, srv)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(netIf.Name + ": ")
				fmt.Println("Hw_addr: " + netIf.Hw_addr.String())
				fmt.Printf("Inet_addr: %+v\n", netIf.Inet_addr)
				fmt.Printf("MTU: %d\n", netIf.MTU)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
	}
}
