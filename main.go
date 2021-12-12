package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/urfave/cli/v2"
)

type AdventOfCode struct {
	Day  int
	Year int16
}

const year int16 = 2021

func main() {
	var day string
	var part string

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "complete a task on the list",
				Action: func(c *cli.Context) error {
					d, err := strconv.Atoi(c.String("day"))
					if err != nil {
						log.Fatal(err)
					}
					createDay(d)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "day",
						Aliases:     []string{"d"},
						Usage:       "day `number`",
						EnvVars:     []string{"DAY"},
						Required:    true,
						Destination: &day,
					},
				},
			},
			{
				Name:  "run",
				Usage: "Run a day challenge.",
				Action: func(c *cli.Context) error {
					d, err := strconv.Atoi(c.String("day"))
					if err != nil {
						log.Fatal(err)
					}
					createDay(d)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "day",
						Aliases:     []string{"d"},
						Usage:       "day `number`",
						EnvVars:     []string{"DAY"},
						Required:    true,
						Destination: &day,
					},
					&cli.StringFlag{
						Name:        "part",
						Aliases:     []string{"p"},
						Usage:       "part `number` of the chalenge",
						EnvVars:     []string{"PART"},
						Required:    true,
						Destination: &part,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createDay(day int) {

	//Create day folder
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	daysFolderPath := fmt.Sprint(path, "/days")
	_, err = os.Stat(daysFolderPath)
	if os.IsNotExist(err) {
		os.Mkdir(daysFolderPath, 0766)
	}

	dayFolderPath := fmt.Sprint(path, "/days/day", day)
	_, err = os.Stat(dayFolderPath)
	if !os.IsNotExist(err) {
		log.Fatalf("Folder already exist")
	} else {
		os.Mkdir(dayFolderPath, 0766)
	}

	//Generate day golang file
	templatePath := fmt.Sprint(path, "/template.txt")
	dayPath := fmt.Sprint(dayFolderPath, "/day", day, ".go")

	d := AdventOfCode{day, year}
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(dayPath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = tmpl.Execute(f, d)
	if err != nil {
		panic(err)
	}

}
