package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	jd, err := readJSONData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app := cli.NewApp()

	app.EnableBashCompletion = true
	app.Name = "komi"
	app.Usage = "A simple command saver"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Use `komi help` to get help")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "Show commands inside a category",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					displayCategories(jd)
					return nil
				}

				catgr := c.Args().First()
				err = displayCommands(jd, catgr, false)
				if err != nil {
					return err
				}

				return nil
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "add",
			Usage: "Add command to a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				err := addCommand(jd, catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "mod",
			Usage: "Modify command in a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				err := modifyCommand(jd, catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "del",
			Usage: "Delete command from a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				err := deleteCommand(jd, catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "addcat",
			Usage: "Add a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify category name")
				}

				catgr := c.Args().First()
				err := addCategory(jd, catgr)
				return err
			},
		},
		{
			Name:  "modcat",
			Usage: "Modify a category name",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				err := modifyCategory(jd, catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "delcat",
			Usage: "Delete a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				err := deleteCategory(jd, catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
		},
		{
			Name:  "export",
			Usage: "Display all data in JSON",
			Action: func(c *cli.Context) error {
				err = displayJSONData(jd)

				return err
			},
		},
		{
			Name:  "search",
			Usage: "Search for a string inside data",
			Action: func(c *cli.Context) error {
				query := strings.Join(c.Args(), " ")
				ignoreCase := c.Bool("ignore-case")
				err := search(jd, query, ignoreCase)

				return err
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "ignore-case, i",
					Usage: "Ignore case for search query",
				},
			},
		},
		{
			Name:  "copy",
			Usage: "Copy a command or its use text",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category")
				}

				catgr := c.Args().First()
				copyUse := c.Bool("copy-use")
				err := copyCommand(jd, catgr, copyUse)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				jd.printCategories()
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "copy-use, u",
					Usage: "Copy Command's Use text",
				},
			},
		},
	}

	app.Run(os.Args)
}
