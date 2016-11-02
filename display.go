package main

import (
	"errors"
	"fmt"

	"github.com/carmark/pseudo-terminal-go/terminal"
)

func extractCategoryList(jd jsonData) []string {
	var categories []string

	if jd.Categories == nil {
		return categories
	}

	for _, c := range jd.Categories {
		categories = append(categories, c.Name)
	}

	return categories
}

// Returns list of commands and true if category was found
// else it returns an empty commands list with false
func extractCommands(jd jsonData, catgr string) ([]command, bool) {
	var commands []command

	if jd.Categories == nil {
		return commands, false
	}

	for _, c := range jd.Categories {
		if c.Name == catgr {
			return c.Commands, true
		}
	}

	return commands, false
}

// Display categories in a square matrix
func displayCategories(jd jsonData) {
	categories := extractCategoryList(jd)
	l := len(categories)
	// Length of category with max length
	var maxCatLen int

	for _, c := range categories {
		if maxCatLen < len(c) {
			maxCatLen = len(c)
		}
	}

	// Calculate required order for square matrix
	n := 1
	for l > n*n {
		n++
	}

	// Display categories
	for i, c := range categories {
		fmt.Print(c)
		for j := 0; j <= maxCatLen-len(c); j++ {
			fmt.Print(" ")
		}
		if (i+1)%n == 0 {
			fmt.Println()
		}
	}

	// Required newline if categories do not form a
	// complete square matrix
	if l%n != 0 {
		fmt.Println()
	}
}

// Display commands and uses in columns
// Return error if category not found
func displayCommands(jd jsonData, catgr string, serialize bool, comNamesOnly bool) error {
	coms, found := extractCommands(jd, catgr)

	if !found {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	// Length of command with max length
	var maxComLen int

	for _, c := range coms {
		if maxComLen < len(c.Name) {
			maxComLen = len(c.Name)
		}
	}

	// Get terminal size to truncate Commands `Use` string
	w, _, err := terminal.GetSize(0)
	if err != nil {
		w = 80
	}

	// Display columns
	for i, c := range coms {
		spaceOffset := maxComLen - len(c.Name) + 2
		useStringWidth := w - maxComLen - 2

		if serialize {
			fmt.Printf("%-2d ", i+1)
			useStringWidth -= 3
		}
		if useStringWidth < 0 {
			useStringWidth = 0
		}
		if comNamesOnly {
			fmt.Printf("%s\n", c.Name)

		} else {
			fmt.Printf("%s%*s%-.*s\n", c.Name, spaceOffset, " ", useStringWidth, c.Use)
		}
	}

	return nil
}
