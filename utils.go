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

type jsonFileError struct {
	err error
	msg string
}

func (e *jsonFileError) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

// Read input in raw mode
func readInput(prompt string, optional bool) (string, error) {
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
	defer t.ReleaseFromStdInOut()

	for {
		fmt.Print(prompt)
		line, err := t.ReadLine()
		if err != nil {
			return "", err
		}

		input := strings.Trim(line, " ")
		if input != "" || optional {
			return input, nil
		}
	}
}
