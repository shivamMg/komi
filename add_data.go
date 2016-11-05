package main

import (
	"errors"
	"fmt"
)

func addCategory(jd jsonData, catgr string) error {
	if _, err := jd.getIndex(catgr); err == nil {
		msg := fmt.Sprintf("Category named `%s` already exists", catgr)
		return errors.New(msg)
	}

	jd.AddCategory(category{Name: catgr})

	err := writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func addCommand(jd jsonData, catgr string) error {
	index, err := jd.getIndex(catgr)
	if err != nil {
		return err
	}

	name, err := readInput("Command Name: ", false)
	if err != nil {
		return err
	}
	use, err := readInput("Use (opt): ", true)
	if err != nil {
		return err
	}

	jd.AddCommand(index, command{Name: name, Use: use})

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
