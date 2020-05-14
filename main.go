package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MihaiBlebea/go-gibli/reader"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{}
	app.Commands = []cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the models",
			Action: func(c *cli.Context) error {
				path, err := os.Getwd()
				if err != nil {
					log.Println(err)
				}

				config, err := newConfig(path)
				if err != nil {
					log.Fatal(err)
				}
				err = GenerateModels(config)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Task completed")
				return nil
			},
		},
		{
			Name:    "definition",
			Aliases: []string{"d"},
			Usage:   "Generate a definition file",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Required: true},
				&cli.StringFlag{Name: "kind", Required: false},
			},
			Action: func(c *cli.Context) error {
				path, err := os.Getwd()
				if err != nil {
					log.Println(err)
				}

				config, err := newConfig(path)
				if err != nil {
					log.Fatal(err)
				}

				kind := c.String("kind")
				if kind == "" {
					kind = "model"
				}

				err = GenerateDefinitionFile(config, c.String("name"), kind)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Task completed")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func connect(user, password, host string, port int, db string) *sql.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, strconv.Itoa(port), db)
	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func newConfig(path string) (config Config, err error) {
	c, err := reader.ReadConfig(fmt.Sprintf("%s/gibli.yaml", path))
	if err != nil {
		return config, err
	}

	client := func() *sql.DB {
		return connect(
			c.Connection.User,
			c.Connection.Password,
			c.Connection.Host,
			c.Connection.Port,
			c.Connection.DB,
		)
	}

	config.DefinitionsPath = c.Definitions
	config.ModelsPath = c.Models
	config.Client = client
	return config, nil
}
