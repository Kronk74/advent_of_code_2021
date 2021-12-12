package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/Kronk74/advent_of_code_2021/days"
	"github.com/Kronk74/advent_of_code_2021/utils/aocg"
	req "github.com/Kronk74/advent_of_code_2021/utils/req"
	"github.com/urfave/cli/v2"
)

func main() {

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
					aocg.CreateDay(d)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "day",
						Aliases:  []string{"d"},
						Usage:    "day `number`",
						EnvVars:  []string{"DAY"},
						Required: true,
					},
				},
			},
			{
				Name:        "get-input",
				Usage:       "Get input of a specific day",
				Description: "It will download input from advent of code and put it into day's folder related to. Don't use it too much to not overload remote server.",
				Action: func(c *cli.Context) error {
					path, err := os.Getwd()
					aocg.Check(err)

					dayFolderPath := fmt.Sprint(path, "/days/day", c.String("day"))
					_, err = os.Stat(dayFolderPath)
					if os.IsNotExist(err) {
						log.Fatalf("Please create day's folder before running this command", err)
					}

					d, err := strconv.Atoi(c.String("day"))
					aocg.Check(err)

					input := req.RequestInput(d, c.String("sessionid"))
					aocg.Check(err)

					inputPath := fmt.Sprint(dayFolderPath, "/input")
					err = os.WriteFile(inputPath, input, 0766)
					aocg.Check(err)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "day",
						Aliases:  []string{"d"},
						Usage:    "day `number`",
						EnvVars:  []string{"DAY"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "sessionid",
						Aliases:  []string{"s"},
						Usage:    "session id for http request",
						EnvVars:  []string{"AOCG_SESSION"},
						Required: true,
					},
				},
			},
			{
				Name:  "run",
				Usage: "Run a day challenge.",
				Action: func(c *cli.Context) error {
					day, err := strconv.Atoi(c.String("day"))
					aocg.Check(err)
					part, err := strconv.Atoi(c.String("part"))
					aocg.Check(err)
					result := days.CallDay(day, part)

					if c.Bool("answer") {
						req.PushAnswer(result, day, part, c.String("sessionid"))
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "day",
						Aliases:  []string{"d"},
						Usage:    "day `number`",
						EnvVars:  []string{"DAY"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "part",
						Aliases:  []string{"p"},
						Usage:    "part `number` of the chalenge",
						EnvVars:  []string{"PART"},
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "answer",
						Aliases:  []string{"a"},
						Usage:    "Push your answer to the verification server.",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "sessionid",
						Aliases:  []string{"s"},
						Usage:    "session id for http request",
						EnvVars:  []string{"AOCG_SESSION"},
						Required: true,
						Hidden:   true,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
