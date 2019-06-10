package main


import (
  "fmt"
  "os"
  
  "github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "complete",
			Aliases: []string{"c"},
			Usage: "complete a task on list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil { fmt.Print("error => ", err) }
}
