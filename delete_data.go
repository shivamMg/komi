package main

import (
	"errors"
	"fmt"
)

func deleteCategory(jd jsonData, catgr string) error {
	index, err := jd.getIndex(catgr)
	if err != nil {
		return err
	}

	prompt := fmt.Sprintf("You sure want to delete `%s`? (y/N): ", catgr)
	choice, err := readInput(prompt, false)
	if err != nil || !(choice[0:1] == "y" || choice[0:1] == "Y") {
		msg := fmt.Sprintf("`%s` was not deleted", catgr)
		return errors.New(msg)
	}

	jd.DeleteCategory(index)

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}

func deleteCommand(jd jsonData, catgr string) error {
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

	fmt.Println(jd.Categories[index].Commands[sno-1].Name)
	choice, err := readInput("You sure want to delete above command? (y/N): ", false)
	if err != nil || !(choice[0:1] == "y" || choice[0:1] == "Y") {
		return errors.New("Command was not deleted")
	}

	jd.DeleteCommand(index, sno-1)

	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
