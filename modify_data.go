package main

import (
	"errors"
	"fmt"
)

func modifyCategory(jd jsonData, catgr string) error {
	index, err := jd.getIndex(catgr)
	if err != nil {
		return err
	}

	input, err := readInput("Enter new category name: ", false)
	if err != nil {
		return err
	}

	if _, err := jd.getIndex(input); err == nil {
		msg := fmt.Sprintf("Category named `%s` already exists", input)
		return errors.New(msg)
	}

	jd.ModifyCategory(index, input)

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func modifyCommand(jd jsonData, catgr string) error {
	index, err := jd.getIndex(catgr)
	if err != nil {
		return err
	}

	err = displayCommands(jd, catgr, true)
	if err != nil {
		return err
	}

	fmt.Println()
	sno, err := readSerialNo(len(jd.Categories[index].Commands))
	if err != nil {
		return err
	}

	fmt.Println("Previous name:", jd.Categories[index].Commands[sno-1].Name)
	name, err := readInput("Enter new name: ", false)
	if err != nil {
		return err
	}
	fmt.Println("Previous use:", jd.Categories[index].Commands[sno-1].Use)
	use, err := readInput("Enter new use (opt): ", true)
	if err != nil {
		return err
	}

	jd.ModifyCommand(index, sno-1, name, use)

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
