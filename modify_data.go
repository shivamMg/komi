package main

import (
	"errors"
	"fmt"
	"strconv"
)

func modifyCategory(catgr string) error {
	jd, err := readJSONData()
	if err != nil {
		return err
	}

	index := jd.getIndex(catgr)

	if index == -1 {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	input := readInput("Enter new category name: ")
	if index := jd.getIndex(input); index != -1 {
		msg := fmt.Sprintf("Category named `%s` already exists\n", input)
		return errors.New(msg)
	}

	jd.Categories[index].Name = input

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func modifyCommand(catgr string) error {
	jd, err := readJSONData()
	if err != nil {
		return err
	}

	index := jd.getIndex(catgr)

	if index == -1 {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	err2 := displayCommands(jd, catgr, true, true)
	if err2 != nil {
		return err2
	}

	fmt.Println()
	input := readInput("Select Serial no: ")
	sno, err2 := strconv.Atoi(input)
	if err2 != nil {
		return err2
	} else if sno < 1 || sno > len(jd.Categories[index].Commands) {
		return errors.New("Invalid Serial no")
	}

	fmt.Println("Previous name:", jd.Categories[index].Commands[sno-1].Name)
	name := readInput("Enter new name: ")
	fmt.Println("Previous use:", jd.Categories[index].Commands[sno-1].Use)
	use := readInput("Enter new use: ")

	jd.Categories[index].Commands[sno-1].Name = name
	jd.Categories[index].Commands[sno-1].Use = use

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
