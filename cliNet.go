package main

import (
  "log"
  "os"
  "fmt"
  "net_cli/request"
  "github.com/urfave/cli"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli_net"
	app.Usage = "cli_net [global options] command [command options] [arguments...]"

	var pr request.Params
	appFlags := []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			Value:       "127.0.0.1",
			Destination: &pr.Server,
		},
		cli.StringFlag{
			Name:        "port",
			Value:       "8080",
			Destination: &pr.Port,
		},
	}

	initVersion := func() {
		ver, err := request.VersionRequest(pr)
		if err != nil {
			log.Fatal(err)
		}
		pr.NetIfVersion = ver.Version
	}

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "Returns all available interfaces on the host and virtual network",
			Flags: appFlags,
			Action: func(c *cli.Context) {
				initVersion()
				listOfNames, err := request.InterfaceListRequest(pr)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println(strings.Join(listOfNames.AllIntr, ","))
			},
		},
		{
			Name:  "show",
			Usage: "Show info about interface whith <name>",
			Flags: appFlags,
			Action: func(c *cli.Context) {
				if c.NArg() == 0 {
					fmt.Println("not enough args")
					return
				}

				initVersion()

				name := c.Args()[0]
				netIf, err := request.InterfaceInfoRequest(name, pr)
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
