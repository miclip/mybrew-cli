package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miclip/mybrewgo/recipe"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mybrewgo"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}
	recipe := recipe.Recipe{}
	log.Printf("Recipe %s", recipe.Name)

	app.Run(os.Args)
}
