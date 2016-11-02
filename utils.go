package main

import (
	"fmt"
	"strings"

	"github.com/carmark/pseudo-terminal-go/terminal"
)

type command struct {
	Name string `json:"name"`
	Use  string `json:"use"`
}

type category struct {
	Name     string    `json:"name"`
	Commands []command `json:"commands"`
}

type jsonData struct {
	Categories []category `json:"categories"`
}

// Returns index of a category name else -1
func (jd *jsonData) getIndex(catgr string) int {
	for i, c := range jd.Categories {
		if c.Name == catgr {
			return i
		}
	}
	return -1
}

type jsonFileError struct {
	err error
	msg string
}

func (e *jsonFileError) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

// Read input in raw mode
func readInput(prompt string) string {
	// Save currect state of terminal
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

	t, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print(prompt)
		line, err := t.ReadLine()
		if err != nil {
			panic(err)
		}

		input := strings.Trim(line, " ")
		if input != "" {
			return input
		}
	}
}

// Prints category names for Bash completion
func printCategories() {
	jd, err := readJSONData()
	if err != nil {
		return
	}

	categories := extractCategoryList(jd)
	for _, c := range categories {
		fmt.Println(c)
	}
}
