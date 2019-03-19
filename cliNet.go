package main

import (
  "log"
  "os"
  "fmt"
  "cli/request"
  "github.com/urfave/cli"
)

func main() {
  var pr request.Params
  app := cli.NewApp()
  app.Name = "cli_net"
  app.Usage = "cli_net [global options] command [command options] [arguments...]"

  appFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "server",
      Value: "127.0.0.1",
      Destination: &pr.Server,
      
    },
    cli.StringFlag{
			Name:  "port",
      Value: "8080",
      Destination: &pr.Port,
    },
    
  }
  
  ver, err := request.VersionRequest(pr)
  if err != nil {
    panic(err)
  }
  pr.NetIfVersion = ver.Version
  
	app.Commands = []cli.Command {
		      {
			            Name:  "list",
                  Usage: "Returns all available interfaces on the host and virtual network", 
                  Flags: appFlags,
                  Action: func(c *cli.Context) {
                    listOfNames, err := request.InterfaceListRequest(pr)
                    if err != nil {
                      fmt.Println(err.Error())
                    } else {
                      var i int
                      for i= 0; i < len(listOfNames.AllIntr)-1;i++ {
                      fmt.Print(listOfNames.AllIntr[i] + ",")
                      }
                      fmt.Println(listOfNames.AllIntr[i])
                    }
                  },
          },
          {
                  Name:  "show",
                  Usage: "Show info about interface whith <name>", 
                  Flags: appFlags,                  
                  Action: func(c *cli.Context) {
                    name := ""
                    if c.NArg() > 0 {
                      name = c.Args()[0]
                    }

                    netIf, err := request.InterfaceInfoRequest(name, pr)
                    if err != nil {
                      fmt.Println(err.Error())
                    } else {
                      fmt.Println(netIf.Name + ": ")
                      fmt.Println("Hw_addr: " + netIf.Hw_addr.String())
                      fmt.Printf("Inet_addr: %+v\n", netIf.Inet_addr)
                      fmt.Printf("MTU: %d\n", netIf.MTU)
                      
                    }
                  },
            },
          
	  }  
  err = app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}