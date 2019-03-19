package main

import (
  "log"
  "os"
  "fmt"
  

  "github.com/urfave/cli"


)

type Params struct {
  NetIfVersion string
  Server string
  Port  string

}

func main() {
  var pr Params
  ver, err := MakeVersionRequest()
  if err != nil {
    //todo add error  
  }
  pr.NetIfVersion = ver.Version
  //fmt.Println(NetIfVersion)
  
  
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

	app.Commands = []cli.Command {
		      {
			            Name:  "list",
                  Usage: "Returns all available interfaces on the host and virtual network", //to check virtnetw
                  Flags: appFlags,
                  
                  Action: func(c *cli.Context) {
                    listOfNames, err := MakeListRequest(pr)
                    
                    if err != nil {
                      fmt.Println("eror1")
                      fmt.Println(err.Error())
                    } else {
                      for i := range listOfNames.AllIntr {
                      fmt.Print(listOfNames.AllIntr[i] + ",")
                      }
                      fmt.Println()
                    }
                  },
          },
    
          {
                  Name:  "show",
                  Usage: "Show info about interface whith <name>", //check spell
                  Flags: appFlags,
                  
                  Action: func(c *cli.Context) {
                    name := ""
                    if c.NArg() > 0 {
                      name = c.Args()[0]
                    }
                    netIf, err := MakeInterfaceRequest(name, pr)
                    
                    if err != nil {
                      fmt.Println(err.Error())
                    } else {
                      fmt.Println(netIf.Name + ": ")
                      fmt.Println("Hw_addr: " + netIf.Hw_addr.String())
                      fmt.Printf("Inet_addr: %+v\n", netIf.Inet_addr)
                      //fmt.Printf("%#v" + netIf.Inet_addr)
                      //fmt.Println("Inet_addr: " + netIf.Inet_addr)
                      //fmt.Printf("%+v", netIf.MTU)
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