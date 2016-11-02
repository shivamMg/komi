package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
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
				jd, err := readJSONData()
				if err != nil {
					return err
				}

				if c.NArg() == 0 {
					displayCategories(jd)
					return nil
				}

				catgr := c.Args().First()
				err2 := displayCommands(jd, catgr, false, false)
				if err2 != nil {
					return err2
				}

				return nil
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
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
				err := addCommand(catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
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
				err := modifyCommand(catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
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
				err := deleteCommand(catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
			},
		},
		{
			Name:  "addcat",
			Usage: "Add a category",
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return errors.New("Please specify a category name")
				}

				catgr := c.Args().First()
				err := addCategory(catgr)
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
				err := modifyCategory(catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
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
				err := deleteCategory(catgr)
				return err
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				printCategories()
			},
		},
	}

	app.Run(os.Args)
}
