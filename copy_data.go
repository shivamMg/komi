package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func copyCommand(jd jsonData, catgr string, copyUse bool) error {
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

	cm := jd.Categories[index].Commands[sno-1]

	var text string
	if copyUse {
		text = cm.Use
	} else {
		text = cm.Name
	}

	err = clipboard.WriteAll(text)
	if err != nil {
		// return errors.New("Command/Use not copied")
		return err
	}

	fmt.Printf("`%s` Copied!\n", text)

	return nil
}
