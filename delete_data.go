package main

import (
	"errors"
	"fmt"
	"strconv"
)

func deleteCategory(catgr string) error {
	jd, err := readJSONData()
	if err != nil {
		return err
	}

	index := jd.getIndex(catgr)

	if index == -1 {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	prompt := fmt.Sprintf("You sure want to delete `%s`? (y/N): ", catgr)
	choice := readInput(prompt)
	if err != nil || !(choice[0:1] == "y" || choice[0:1] == "Y") {
		msg := fmt.Sprintf("`%s` was not deleted\n", catgr)
		return errors.New(msg)
	}

	l := len(jd.Categories)
	// Swap with last category
	jd.Categories[index] = jd.Categories[l-1]
	// Slice off last category
	jd.Categories = jd.Categories[:l-1]

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func deleteCommand(catgr string) error {
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
		return errors.New("Invalid Serial no\n")
	}

	fmt.Println(jd.Categories[index].Commands[sno-1].Name)
	choice := readInput("You sure want to delete above command? (y/N): ")
	if err != nil || !(choice[0:1] == "y" || choice[0:1] == "Y") {
		return errors.New("Command was not deleted\n")
	}

	l := len(jd.Categories[index].Commands)
	// Swap with last category
	jd.Categories[index].Commands[sno-1] = jd.Categories[index].Commands[l-1]
	// Slice off last category
	jd.Categories[index].Commands = jd.Categories[index].Commands[:l-1]

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil

}
