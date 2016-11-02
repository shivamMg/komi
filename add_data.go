package main

import (
	"errors"
	"fmt"
)

func addCategory(catgr string) error {
	c := category{Name: catgr}

	jd, err := readJSONData()
	if err != nil {
		return err
	}

	jd.Categories = append(jd.Categories, c)
	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func addCommand(catgr string) error {
	jd, err := readJSONData()
	if err != nil {
		return err
	}

	index := jd.getIndex(catgr)
	if index == -1 {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	name := readInput("Command Name: ")
	use := readInput("Use: ")

	com := command{Name: name, Use: use}

	jd.Categories[index].Commands = append(jd.Categories[index].Commands, com)
	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
