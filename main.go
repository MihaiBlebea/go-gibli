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

// 1. Get the arguments
func main() {
	// bundle.Bundle()
	// config := Config{ModelDefinitionPath: "./definitions", ModelFilesPath: "./foo", Client: connect}

	// err := GenerateModels(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
		Commands: []cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build the models",
				Action: func(c *cli.Context) error {
					path, err := os.Getwd()
					if err != nil {
						log.Println(err)
					}

					config, err := reader.ReadConfig(fmt.Sprintf("%s/gibli.yaml", path))
					if err != nil {
						log.Println(err)
					}

					client := func() *sql.DB {
						return connect(
							config.Connection.User,
							config.Connection.Password,
							config.Connection.Host,
							config.Connection.Port,
							config.Connection.DB,
						)
					}

					cf := Config{
						ModelDefinitionPath: config.Definitions,
						ModelFilesPath:      config.Models,
						Client:              client,
					}

					err = GenerateModels(cf)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println(config)
					fmt.Println("Task completed")
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("Task added")
					return nil
				},
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
